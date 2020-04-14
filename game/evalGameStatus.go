package game

import (
	game "github.hc.ag/jsuarez/chopsticks/game/operationsImplementation"
)

func IsWin(players []game.PlayerI) bool {
	for _, player := range players {
		if player.GetPlayer().LeftHand == 0 && player.GetPlayer().RightHand == 0 {
			return true
		}
	}
	return false
}
