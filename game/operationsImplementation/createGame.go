package game

import (
	"math/rand"
	"time"

	"github.hc.ag/jsuarez/chopsticks/config"
)

//Player is an structure that represent each game participant
type Player struct {
	LeftHand  int
	RightHand int
}

//PlayerI is the interface for the different implementations of the player and its actions
type PlayerI interface {
	GetPlayer() *Player
	playSplit()
	playAttack(*Player, int, string)
}

//CreateGame initialize the Number of players defined in the configutarion to the initial state
func CreateGame(config *config.Configuration) []PlayerI {
	players := make([]PlayerI, config.Players)
	for i := 0; i < config.Players; i++ {
		//players[i] = &PlayerOperationsImpl{&Player{1, 1}}
		players[i] = &PlayerOperationsImpl{&Player{1, 1}}
	}
	return players
}

//StartGame comence the first turn of the game with a random selected player
func StartGame(config *config.Configuration) int {
	rand.Seed(time.Now().UnixNano())
	playerTurn := rand.Intn(config.Players)
	return playerTurn
}
