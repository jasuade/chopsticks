package main

import (
	"log"

	"github.hc.ag/jsuarez/chopsticks/config"
)

const DEFAULT_GAME_CONFIGURATION_PATH = "game.config"

func main() {
	_, err := config.OpenConfig(DEFAULT_GAME_CONFIGURATION_PATH)
	if err != nil {
		log.Fatal(err)
	}
	//Read the configuration file
	//Set the game
	//Start the game
}
