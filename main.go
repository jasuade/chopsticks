package main

import (
	"log"

	"github.hc.ag/jsuarez/chopsticks/config"
	"github.hc.ag/jsuarez/chopsticks/game"
	"github.hc.ag/jsuarez/chopsticks/ui"
)

const DEFAULT_GAME_CONFIGURATION_PATH = "game.config"

func main() {
	//Read configuration file
	config, err := config.OpenConfig(DEFAULT_GAME_CONFIGURATION_PATH)
	if err != nil {
		log.Fatal(err)
	}
	//Set the game
	players := make([]game.Player, config.Players)
	players = game.CreateGame(config)
	//Start the game
	game.StartGame(players)
	ui.DisplayStatus(players)

}
