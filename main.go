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

	config, err := config.OpenConfig(DEFAULT_GAME_CONFIGURATION_PATH)
	if err != nil {
		log.Fatal(err)
	}

	players := game.CreateGame(config)

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

}

func GameLoop(controller *sdlUi.SDLController, players []game.PlayerI, playerTurn int, config *config.Configuration) error {
	gameState := &ui.GameState{players, playerTurn, fmt.Sprintf(ui.Messages["MSG_START"], config.Players, config.Mode, config.Difficulty), 0, 0}

	renderGameState := func() {
		controller.Renderer.Clear()
		controller.PrintBackground()
		controller.PrintTitle()
		controller.PrintMessage(gameState.Msg)
		controller.PrintPlayers(players)
		controller.Renderer.Present()

	}
	renderGameState()
	go func() {
		time.Sleep(2 * time.Second)
		nextRound(gameState)
	}()
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch typedEvent := event.(type) {
			case *sdl.QuitEvent:
				controller.EndGame()
				return nil
			case *sdl.KeyboardEvent:
				if typedEvent.Type == sdl.KEYDOWN {
					updateGameState(gameState, typedEvent.Keysym.Sym)
				}
				continue
			}
		}
		renderGameState()
	}
}

func msgAction(gameState *ui.GameState) {
	gameState.TurnState = ui.ATTACK_OR_SPLIT
	go func() {
		time.Sleep(2 * time.Second)
		gameState.Msg = fmt.Sprintf(ui.Messages["MSG_ACTION"])
	}()
}

func nextRound(gameState *ui.GameState) {
	gameState.PlayerTurn = (gameState.PlayerTurn + 1) % 2
	go func() {
		time.Sleep(2 * time.Second)
		gameState.Msg = fmt.Sprintf(ui.Messages["MSG_TURN"], gameState.PlayerTurn)
		msgAction(gameState)
	}()
}

func invalidHand(gameState *ui.GameState, nextState byte, message string) {
	gameState.Msg = fmt.Sprintf(ui.Messages["ERR_INVALID_ATTACK"])
	go func() {
		time.Sleep(2 * time.Second)
		gameState.Msg = fmt.Sprintf(message)
	}()
	gameState.TurnState = nextState

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
						msgAction(gameState)
						return
					}
					break
				}
			default:
				{
					gameState.Msg = fmt.Sprintf(ui.Messages["ERR_INVALID_ACTION"])
					msgAction(gameState)
					return

				}
			}
			nextRound(gameState)
			return
		}
	case ui.LEFT_OR_RIGHT_ATTACKER:
		{
			switch key {
			case sdl.K_l:
				if currentPlayer.GetPlayer().LeftHand > 0 {
					gameState.AttackPower = currentPlayer.GetPlayer().LeftHand
					break
				}
				invalidHand(gameState, ui.LEFT_OR_RIGHT_ATTACKER, ui.Messages["MSG_HAND_TO_ATTACK_WITH"])
				return
			case sdl.K_r:
				if currentPlayer.GetPlayer().RightHand > 0 {
					gameState.AttackPower = currentPlayer.GetPlayer().RightHand
					break
				}
				invalidHand(gameState, ui.LEFT_OR_RIGHT_ATTACKER, ui.Messages["MSG_HAND_TO_ATTACK_WITH"])
				return
			default:
				{
					invalidHand(gameState, ui.LEFT_OR_RIGHT_ATTACKER, ui.Messages["MSG_HAND_TO_ATTACK_WITH"])
					return

				}
			}
			gameState.Msg = ui.Messages["MSG_HAND_TO_ATTACK_TO"]
			gameState.TurnState = ui.LEFT_OR_RIGHT_DEFENDER
			return
		}
	case ui.LEFT_OR_RIGHT_DEFENDER:
		{
			var receiverHand string
			switch key {
			case sdl.K_l:
				if oponentPlayer.LeftHand > 0 {
					receiverHand = "l"
					break
				}
				invalidHand(gameState, ui.LEFT_OR_RIGHT_DEFENDER, ui.Messages["MSG_HAND_TO_ATTACK_TO"])
				return
			case sdl.K_r:
				if oponentPlayer.RightHand > 0 {
					receiverHand = "r"
					break
				}
				invalidHand(gameState, ui.LEFT_OR_RIGHT_DEFENDER, ui.Messages["MSG_HAND_TO_ATTACK_TO"])
				return
			default:
				{
					invalidHand(gameState, ui.LEFT_OR_RIGHT_DEFENDER, ui.Messages["MSG_HAND_TO_ATTACK_TO"])
					return
				}
			}
			if err := gameState.Players[gameState.PlayerTurn].PlayAttack(oponentPlayer, gameState.AttackPower, receiverHand); err != nil {
				gameState.Msg = fmt.Sprintf(ui.Messages["ERR_INVALID_ATTACK"])
				msgAction(gameState)
				return
			}

			if eval.IsWin(gameState.Players) {
				gameState.TurnState = ui.IS_WIN
				gameState.Msg = fmt.Sprintf(ui.Messages["MSG_WIN"], gameState.PlayerTurn)
				return
			}
			nextRound(gameState)
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
			nextRound(gameState)
			return
		}
	}
}
