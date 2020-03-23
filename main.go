package main

import (
	"log"

	"github.hc.ag/jsuarez/chopsticks/config"
)

const DEFAULT_GAME_CONFIGURATION_PATH = "game.config"

func main() {
	//Read configuration file
	config, err := config.OpenConfig(DEFAULT_GAME_CONFIGURATION_PATH)
	if err != nil {
		log.Fatal(err)
	}
	//Set the game
	createGame(config)
	//Start the game
}
