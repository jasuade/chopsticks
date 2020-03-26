package game

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

//PlayTurn execute the action of a player in a given moment, it uses the stdio to read from the console
func PlayTurn(players []Player, playerTurn int) string {
	return playTurnFromReader(players, playerTurn, os.Stdin)
}

func playTurnFromReader(players []Player, playerTurn int, r io.Reader) string {
	reader := bufio.NewReader(r)

	fmt.Print("Are you goint to attack(a) or to split(s):\n")
	action, _ := reader.ReadString('\n')
	switch strings.TrimSpace(action) {
	case "a":
		fmt.Println("You choose to attack")
		action = "attack"
		//attack()
	case "s":
		fmt.Println("You choose to split")
		action = "split"
		players[playerTurn].playSplit()
	default:
		action = "Invalid action"
		fmt.Println(action)
	}
	return action
}

func playAttack() {

}

//Should receive a player with an status and return the same player with different status
func (player *Player) playSplit() {
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
			player.containsNumberFour(0)
			return
		}
		if player.RightHand > player.LeftHand {
			player.RightHand--
			player.LeftHand++
			return
		}
		player.RightHand++
		player.LeftHand--
		return
	} else if player.RightHand == player.LeftHand {
		player.RightHand--
		player.LeftHand++
		return
	} else if int(math.Abs(float64(player.LeftHand-player.RightHand))) == 1 {
		if containsNumber(player, 2) && containsNumber(player, 3) {
			player.LeftHand = 1
			player.RightHand = 4
			return
		}
		fmt.Printf("That is not a meaninful split, as %v will be {%d,%d} \n", player, player.RightHand, player.LeftHand)
		return

	} else if player.RightHand > player.LeftHand {
		player.RightHand--
		player.LeftHand++
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

func (player *Player) containsNumberFour(num int) {
	fmt.Println("How many chopsticks you want to transfere?")
	fmt.Scan(&num)
	if num == 1 {
		if player.RightHand > player.LeftHand {
			player.RightHand--
			player.LeftHand++
			return
		}
		player.RightHand++
		player.LeftHand--
		return
	}
	player.RightHand = 2
	player.LeftHand = 2
	return

}
