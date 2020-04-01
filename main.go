package main

import (
	"fmt"
	"log"

	"github.hc.ag/jsuarez/chopsticks/config"
	eval "github.hc.ag/jsuarez/chopsticks/game"
	game "github.hc.ag/jsuarez/chopsticks/game/operationsImplementation"
	ui "github.hc.ag/jsuarez/chopsticks/ui"
)

const DEFAULT_GAME_CONFIGURATION_PATH = "game.config"

func main() {
	//Read configuration file
	config, err := config.OpenConfig(DEFAULT_GAME_CONFIGURATION_PATH)
	if err != nil {
		log.Fatal(err)
	}
	//Set the game
	players := game.CreateGame(config)

	//Start the game
	fmt.Printf("Start the game with %d payers in mode <%s> and level <%s> \n", config.Players, config.Mode, config.Difficulty)
	playerTurn := game.StartGame(config)
	ui.DisplayStatus(players)

	//Turn loop
	for eval.IsWin(players) {
		fmt.Printf("Player %d is your turn\n", playerTurn)
		game.PlayTurn(players, playerTurn)
		game.CheckHandsStatus(players)
		ui.DisplayStatus(players)
		playerTurn = (playerTurn + 1) % 2
	}

}
