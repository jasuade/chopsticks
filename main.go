package main

import (
	"fmt"
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.hc.ag/jsuarez/chopsticks/config"
	eval "github.hc.ag/jsuarez/chopsticks/game"
	game "github.hc.ag/jsuarez/chopsticks/game/operationsImplementation"
	"github.hc.ag/jsuarez/chopsticks/ui"
	sdlUi "github.hc.ag/jsuarez/chopsticks/ui/sdlImplementation"
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
	log.Printf("Start the game with %d payers in mode <%s> and level <%s> \n", config.Players, config.Mode, config.Difficulty)
	playerTurn := game.StartGame(config)

	//ui.DisplayStatus(players)
	//SDL UI
	sdlController, err := sdlUi.InitSDL()
	if err != nil {
		log.Fatal(err)
	}

	if err := GameLoop(sdlController, players, playerTurn, config); err != nil {
		log.Fatal(err)
	}

	//Turn loop
	// for eval.IsWin(players) {
	// 	fmt.Printf("Player %d is your turn\n", playerTurn)
	// 	err = game.PlayTurn(players, playerTurn)
	// 	if err == nil {
	// 		game.CheckHandsStatus(players)
	// 		cmd.DisplayStatus(players)
	// 		playerTurn = (playerTurn + 1) % 2
	// 	} else {
	// 		fmt.Println(err)
	// 	}

	// }

}

func GameLoop(c *sdlUi.SDLController, players []game.PlayerI, playerTurn int, config *config.Configuration) error {
	gameState := &ui.GameState{players, playerTurn, fmt.Sprintf(ui.Messages["MSG_START"], config.Players, config.Mode, config.Difficulty), 0, ui.ATTACK_OR_SPLIT}

	renderGameState := func() {
		c.Renderer.Clear()
		c.PrintBackground()
		c.PrintTitle()
		c.PrintMessage(gameState.Msg)
		c.PrintPlayers(players)
		c.Renderer.Present()

	}
	renderGameState()
	go func() {
		time.Sleep(3 * time.Second)
		gameState.Msg = fmt.Sprintf(ui.Messages["MSG_TURN"], playerTurn)
		go func() {
			time.Sleep(3 * time.Second)
			gameState.Msg = fmt.Sprintf(ui.Messages["MSG_ACTION"])
		}()
	}()
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch typedEvent := event.(type) {
			case *sdl.QuitEvent:
				c.EndGame()
				return nil
			case *sdl.KeyboardEvent:
				if typedEvent.Type == sdl.KEYDOWN {
					updateGameState(gameState, typedEvent.Keysym.Sym)
				}
			}
		}
		renderGameState()
	}
}

func updateGameState(gameState *ui.GameState, key sdl.Keycode) {
	currentPlayer := gameState.Players[gameState.PlayerTurn]
	oponentPlayer := gameState.Players[(gameState.PlayerTurn+1)%2].GetPlayer()

	switch gameState.TurnState {
	case ui.ATTACK_OR_SPLIT:
		{
			switch key {
			case sdl.K_a:
				gameState.Msg = ui.Messages["MSG_HAND_TO_ATTACK_WITH"]
				gameState.TurnState = ui.LEFT_OR_RIGHT_ATTACKER
				return
			case sdl.K_s:
				if game.ContainsNumber(currentPlayer.GetPlayer(), 0) {
					if game.ContainsNumber(currentPlayer.GetPlayer(), 4) {
						gameState.Msg = fmt.Sprintf(ui.Messages["MSG_SPLIT"])
						gameState.TurnState = ui.ONE_OR_TWO
						return
					}
				} else {
					if err := gameState.Players[gameState.PlayerTurn].PlaySplit(); err != nil {
						gameState.Msg = fmt.Sprintf(ui.Messages["ERR_SPLIT"])
						go func() {
							time.Sleep(3 * time.Second)
							gameState.Msg = fmt.Sprintf(ui.Messages["MSG_ACTION"])
						}()
						gameState.TurnState = ui.ATTACK_OR_SPLIT
						return
					}
				}
			default:
				{
					gameState.Msg = fmt.Sprintf(ui.Messages["ERR_INVALID_ACTION"])
					go func() {
						time.Sleep(3 * time.Second)
						gameState.Msg = fmt.Sprintf(ui.Messages["MSG_ACTION"])
					}()
					gameState.TurnState = ui.ATTACK_OR_SPLIT
					return

				}
			}
			gameState.PlayerTurn = (gameState.PlayerTurn + 1) % 2
			gameState.TurnState = ui.ATTACK_OR_SPLIT
			gameState.Msg = fmt.Sprintf(ui.Messages["MSG_TURN"], gameState.PlayerTurn)
			go func() {
				time.Sleep(3 * time.Second)
				gameState.Msg = fmt.Sprintf(ui.Messages["MSG_ACTION"])
			}()
			return
		}
	case ui.LEFT_OR_RIGHT_ATTACKER:
		{
			switch key {
			case sdl.K_l:
				gameState.AttackPower = currentPlayer.GetPlayer().LeftHand
			case sdl.K_r:
				gameState.AttackPower = currentPlayer.GetPlayer().RightHand
			default:
				{
					gameState.Msg = fmt.Sprintf(ui.Messages["ERR_INVALID_ATTACK"])
					go func() {
						time.Sleep(3 * time.Second)
						gameState.Msg = fmt.Sprintf(ui.Messages["MSG_HAND_TO_ATTACK_WITH"])
					}()
					gameState.TurnState = ui.LEFT_OR_RIGHT_ATTACKER
					return

				}
			}
			gameState.Msg = ui.Messages["MSG_HAND_TO_ATTACK_TO"]
			gameState.TurnState = ui.LEFT_OR_RIGHT_DEFENDER
		}
	case ui.LEFT_OR_RIGHT_DEFENDER:
		{
			var receiverHand string
			switch key {
			case sdl.K_l:
				receiverHand = "l"
			case sdl.K_r:
				receiverHand = "r"
			default:
				{
					gameState.Msg = fmt.Sprintf(ui.Messages["ERR_INVALID_ATTACK"])
					go func() {
						time.Sleep(3 * time.Second)
						gameState.Msg = fmt.Sprintf(ui.Messages["MSG_HAND_TO_ATTACK_TO"])
					}()
					gameState.TurnState = ui.LEFT_OR_RIGHT_DEFENDER
					return

				}
			}
			if err := gameState.Players[gameState.PlayerTurn].PlayAttack(oponentPlayer, gameState.AttackPower, receiverHand); err != nil {
				gameState.Msg = fmt.Sprintf(ui.Messages["ERR_INVALID_ATTACK"])
				go func() {
					time.Sleep(3 * time.Second)
					gameState.Msg = fmt.Sprintf(ui.Messages["MSG_ACTION"])
				}()
				gameState.TurnState = ui.ATTACK_OR_SPLIT
				return
			}
			if eval.IsWin(gameState.Players) {
				gameState.TurnState = ui.IS_WIN
				return
			}
			gameState.PlayerTurn = (gameState.PlayerTurn + 1) % 2
			gameState.TurnState = ui.ATTACK_OR_SPLIT
			gameState.Msg = fmt.Sprintf(ui.Messages["MSG_TURN"], gameState.PlayerTurn)
			go func() {
				time.Sleep(3 * time.Second)
				gameState.Msg = fmt.Sprintf(ui.Messages["MSG_ACTION"])
			}()
			return
		}
	case ui.ONE_OR_TWO:
		{
			switch key {
			case sdl.K_1:
				currentPlayer.GetPlayer().HigherToLower()
			case sdl.K_2:
				currentPlayer.GetPlayer().RightHand = 2
				currentPlayer.GetPlayer().LeftHand = 2
			}
			gameState.PlayerTurn = (gameState.PlayerTurn + 1) % 2
			gameState.TurnState = ui.ATTACK_OR_SPLIT
			gameState.Msg = fmt.Sprintf(ui.Messages["MSG_TURN"], gameState.PlayerTurn)
			go func() {
				time.Sleep(3 * time.Second)
				gameState.Msg = fmt.Sprintf(ui.Messages["MSG_ACTION"])
			}()
			return
		}
	case ui.IS_WIN:
		{
			go func() {
				time.Sleep(3 * time.Second)
				gameState.Msg = fmt.Sprintf(ui.Messages["MSG_WIN"], gameState.PlayerTurn)
			}()
			return

		}

	}
}
