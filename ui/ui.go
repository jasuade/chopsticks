//to create a more advance GUI with Electron: https://medium.com/benchkram/tutorial-adding-a-gui-to-golang-6aca601e277d
package ui

import (
	"fmt"

	"github.hc.ag/jsuarez/chopsticks/game"
)

const EMOJI_FINGER string = "🥢"
const EMOJI_WIN_GAME string = "🏆"
const EMOJI_LOST_HAND string = "💢"
const EMOJI_LOST_GAME string = "💩"

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
