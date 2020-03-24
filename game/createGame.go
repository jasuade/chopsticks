package game

import (
	"math/rand"
	"time"

	"github.hc.ag/jsuarez/chopsticks/config"
)

//Player is an structure that represent each of the game participants
type Player struct {
	RightHand int
	LeftHand  int
}

//CreateGame initialize the Number of players defined in the configutarion to the init state
func CreateGame(config *config.Configuration) []Player {
	players := make([]Player, config.Players)
	for i := 0; i < config.Players; i++ {
		players[i] = Player{1, 1}
	}
	return players
}

//StartGame comence the first turn of the a random selected player
func StartGame(players []Player, config *config.Configuration) int {
	rand.Seed(time.Now().UnixNano())
	playerTurn := rand.Intn(config.Players)
	for {
		//PlayTurn(players, playerTurn)
		return playerTurn
		//EvalStatus(players)
	}
}

// func PlayTurn(players []Player, playerTurn int) {
// 	DisplayStatus(players)
// 	// reader := bufio.NewReader(os.Stdin)
// 	// fmt.Print("Enter your city: ")
// 	// city, _ := reader.ReadString('\n')
// 	// fmt.Print("You live in " + city)
// }
