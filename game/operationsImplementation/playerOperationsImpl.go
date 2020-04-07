package game

import (
	"bufio"
	"errors"
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
func chooseAttack(players []PlayerI, playerTurn int, r io.Reader) error {
	atackingPlayer := players[playerTurn].GetPlayer()
	oponentPlayer := players[(playerTurn+1)%2].GetPlayer() // TODO: Modify for multiple players
	// TODO: Modify for multiple players
	// if len(players) > 2 {
	// 	fmt.Printf("Which player %v do you want to attack:\n", players)
	// }

	reader := bufio.NewReader(r)
	fmt.Printf("With which hand (left(l) or right(r))do you want to attack:\n")
	attackerhand, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Printf("And which hand (left(l) or right(r))do you want to attack:\n")
	receiverHand, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	switch strings.TrimSpace(attackerhand) {
	case "l":
		if atackingPlayer.LeftHand < 5 && atackingPlayer.LeftHand > 0 {
			err = players[playerTurn].playAttack(oponentPlayer, atackingPlayer.LeftHand, receiverHand)
			return err
		}
	case "r":
		if atackingPlayer.RightHand < 5 && atackingPlayer.RightHand > 0 {
			err = players[playerTurn].playAttack(oponentPlayer, atackingPlayer.RightHand, receiverHand)
			return err
		}
	}
	err = errors.New("Invalid attack, the hand does not exists")
	return err
}

func (poi *PlayerOperationsImpl) playAttack(oponentPlayer *Player, num int, receiverHand string) error {
	switch strings.TrimSpace(receiverHand) {
	case "l":
		if oponentPlayer.LeftHand < 5 && oponentPlayer.LeftHand > 0 {
			oponentPlayer.LeftHand += num
			return nil
		}
		oponentPlayer.LeftHand = 0
	case "r":
		if oponentPlayer.RightHand < 5 && oponentPlayer.RightHand > 0 {
			oponentPlayer.RightHand += num
			return nil
		}
		oponentPlayer.RightHand = 0
	}

	err := errors.New("Invalid attack, the hand is not alive")
	return err
}

func (poi *PlayerOperationsImpl) playSplit() error {
	player := poi.Player
	if player.LeftHand <= 1 && player.RightHand <= 1 {
		err := errors.New("Unable to slpit, not enough chopsticks, you cannot kill a hand")
		return err
	}
	if player.LeftHand == 4 && player.RightHand == 4 {
		err := errors.New("Unable to slpit, too many chopsticks, you cannot kill a hand")
		return err
	}
	if containsNumber(player, 0) {
		if containsNumber(player, 4) {
			player.splitWithZeroFour(0)
			return nil
		}
		player.higherToLower()
		return nil
	} else if int(math.Abs(float64(player.LeftHand-player.RightHand))) == 1 {
		if containsNumber(player, 2) && containsNumber(player, 3) {
			player.LeftHand = 1
			player.RightHand = 4
			return nil
		}
		err := errors.New("That is not a meaninful split")
		return err

	} else if player.RightHand >= player.LeftHand {
		player.higherToLower()
		return nil
	}
	player.RightHand++
	player.LeftHand--
	return nil
}

func containsNumber(player *Player, i int) bool {
	if player.LeftHand == i || player.RightHand == i {
		return true
	}
	return false
}

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

//CheckHandsStatus is used in Cut-off mode yo equals hands with value >= 5 to value 0
func CheckHandsStatus(players []PlayerI) {
	for _, player := range players {
		if player.GetPlayer().RightHand >= 5 {
			player.GetPlayer().RightHand = 0
		}

		if player.GetPlayer().LeftHand >= 5 {
			player.GetPlayer().LeftHand = 0
		}
	}

}
