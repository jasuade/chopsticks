//to create a more advance GUI with Electron: https://medium.com/benchkram/tutorial-adding-a-gui-to-golang-6aca601e277d
package ui

import (
	"fmt"

	"github.hc.ag/jsuarez/chopsticks/game"
)

//EMOJI_FINGER chopsticks represent a finger in a 5-fingers hand
const EMOJI_FINGER string = "🥢"

//EMOJI_WIN_GAME is displayed when one of the players wins the game
const EMOJI_WIN_GAME string = "🏆"

//EMOJI_LOST_HAND respresent a lost hand (you reach 5 fingers/chopsticks), not a lost game
const EMOJI_LOST_HAND string = "💢"

//EMOJI_LOST_GAME is displayed when one of the players loses the game
const EMOJI_LOST_GAME string = "💩"

//DisplayStatus print on the console the status of the game in a given moment
func DisplayStatus(players []game.PlayerI) error {
	for i, playerI := range players {
		player := playerI.GetPlayer()
		fmt.Printf("P%d", i)
		if player.LeftHand == 5 {
			fmt.Print(" [  " + EMOJI_LOST_HAND + "  ] ")
		} else {
			fmt.Print(" [")
			for i := 0; i < player.LeftHand; i++ {
				fmt.Print(EMOJI_FINGER)
			}
			fmt.Print("] ")
		}

		if player.RightHand == 5 {
			fmt.Print(" [  " + EMOJI_LOST_HAND + "  ] ")
		} else {
			fmt.Print(" [")
			for i := 0; i < player.RightHand; i++ {
				fmt.Print(EMOJI_FINGER)
			}
			fmt.Print("] ")
		}
		fmt.Println("\n___________")
	}

	return nil
}
