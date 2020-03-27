package game

import (
	"fmt"
	"math"
)

type PlayerOperationsImpl struct {
	player *Player
}

func (poi *PlayerOperationsImpl) GetPlayer() *Player {
	return poi.player
}

func (poi *PlayerOperationsImpl) playAttack() {
}

//Should receive a player with an status and return the same player with different status
func (poi *PlayerOperationsImpl) playSplit() {
	player := poi.player
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
