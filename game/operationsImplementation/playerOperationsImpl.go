package game

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

//PlayerOperationsImpl is the current implementation of the interface PlayerI, implementation done with math operations
type PlayerOperationsImpl struct {
	Player *Player
}

//GetPlayer function restuns the Current implementation of the Player
func (poi *PlayerOperationsImpl) GetPlayer() *Player {
	return poi.Player
}

//Should receive all players with its status, execute attack from a player and return the new status
func chooseAttack(players []PlayerI, playerTurn int, r io.Reader) {
	atackingPlayer := players[playerTurn].GetPlayer()
	defendingPlayer := players[(playerTurn+1)%2].GetPlayer() // TODO: Modify for multiple players
	// TODO: Modify for multiple players
	// if len(players) > 2 {
	// 	fmt.Printf("Which player %v do you want to attack:\n", players)
	// }
	reader := bufio.NewReader(r)
	fmt.Printf("With which hand (left(l) or right(r))do you want to attack:\n")
	attackerhand, _ := reader.ReadString('\n')

	fmt.Printf("And which hand (left(l) or right(r))do you want to attack:\n")

	receiverHand, _ := reader.ReadString('\n')
	switch strings.TrimSpace(attackerhand) {
	case "l":
		players[playerTurn].playAttack(defendingPlayer, atackingPlayer.LeftHand, receiverHand)
	case "r":
		players[playerTurn].playAttack(defendingPlayer, atackingPlayer.RightHand, receiverHand)
	}
}

func (poi *PlayerOperationsImpl) playAttack(defendingPlayer *Player, num int, receiverHand string) {
	switch strings.TrimSpace(receiverHand) {
	case "l":
		defendingPlayer.LeftHand += num
	case "r":
		defendingPlayer.RightHand += num
	}
}

//Should receive a player with an status and return the same player with different status
func (poi *PlayerOperationsImpl) playSplit() {
	player := poi.Player
	if player.LeftHand <= 1 && player.RightHand <= 1 {
		fmt.Println("Unable to slpit, not enough chopsticks, you cannot kill a hand")
		return
	}
	if player.LeftHand == 4 && player.RightHand == 4 {
		fmt.Println("Unable to slpit, too many chopsticks, you cannot kill a hand")
		return
	}
	if containsNumber(player, 0) {
		if containsNumber(player, 4) {
			player.splitWithZeroFour(0)
			return
		}
		player.higherToLower()
		return
	} else if int(math.Abs(float64(player.LeftHand-player.RightHand))) == 1 {
		if containsNumber(player, 2) && containsNumber(player, 3) {
			player.LeftHand = 1
			player.RightHand = 4
			return
		}
		fmt.Printf("That is not a meaninful split, as %v will be {%d,%d} \n", player, player.RightHand, player.LeftHand)
		return
	} else if player.RightHand >= player.LeftHand {
		player.higherToLower()
		return
	}
	player.RightHand++
	player.LeftHand--
}

//Receives a player and a number return true if the number is in the arra, false otherwise
func containsNumber(player *Player, i int) bool {
	if player.LeftHand == i || player.RightHand == i {
		return true
	}
	return false
}

//Increase the value of the lower number in 1 and decreases the higher in 1
func (player *Player) higherToLower() {
	if player.RightHand >= player.LeftHand {
		player.RightHand--
		player.LeftHand++
		return
	}
	player.RightHand++
	player.LeftHand--
	return
}

func (player *Player) splitWithZeroFour(num int) {
	fmt.Println("How many chopsticks you want to transfere?")
	fmt.Scan(&num)
	if num == 1 {
		player.higherToLower()
		return
	}
	player.RightHand = 2
	player.LeftHand = 2
	return

}
