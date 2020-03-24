package game

import (
	"github.hc.ag/jsuarez/chopsticks/config"
)

type Player struct {
	RightHand int
	LeftHand  int
}

func createGame(config *config.Configuration) []Player {
	players := make([]Player, config.Players)
	for i := 0; i < config.Players; i++ {
		players[i] = Player{1, 1}
	}
	return players
}
