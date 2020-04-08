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

	players := game.CreateGame(config)

	//Start the game
	fmt.Printf("Start the game with %d payers in mode <%s> and level <%s> \n", config.Players, config.Mode, config.Difficulty)
	playerTurn := game.StartGame(config)

	//ui.DisplayStatus(players)
	//SDL UI
	w, r, resources, err := ui.InitSDL()
	if err != nil {
		log.Fatal(err)
	}
	if err := ui.PrintPlayers(players, r, resources); err != nil {
		log.Fatal(err)
	}

	if err := ui.GameLoop(w, r); err != nil {
		log.Fatal(err)
	}

	//Turn loop
	for eval.IsWin(players) {
		fmt.Printf("Player %d is your turn\n", playerTurn)
		err = game.PlayTurn(players, playerTurn)
		if err == nil {
			game.CheckHandsStatus(players)
			ui.DisplayStatus(players)
			playerTurn = (playerTurn + 1) % 2
		} else {
			fmt.Println(err)
		}

	}

}
