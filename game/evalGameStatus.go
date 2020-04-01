package game

import (
	"fmt"

	game "github.hc.ag/jsuarez/chopsticks/game/operationsImplementation"
	ui "github.hc.ag/jsuarez/chopsticks/ui"
)

func IsWin(players []game.PlayerI) bool {
	p0 := players[0].GetPlayer()
	p1 := players[1].GetPlayer()
	if p0.LeftHand == 0 && p0.RightHand == 0 {
		fmt.Println("Player P1 WINS the game")
		fmt.Println(ui.EMOJI_WIN_GAME)
		fmt.Println("Player P0 LOSES the game")
		fmt.Println(ui.EMOJI_LOST_GAME)
		return false
	}
	if p1.LeftHand == 0 && p1.RightHand == 0 {
		fmt.Println("Player P0 WINS the game")
		fmt.Println(ui.EMOJI_WIN_GAME)
		fmt.Println("Player P1 LOSES the game")
		fmt.Println(ui.EMOJI_LOST_GAME)
		return false
	}
	return true
}
