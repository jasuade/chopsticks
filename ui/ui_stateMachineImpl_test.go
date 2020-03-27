package ui

import (
	"testing"

	game "github.hc.ag/jsuarez/chopsticks/game/stateMachineImplementation"
)

func TestUIDisplaying(t *testing.T) {
	t.Run("When call DisplayStatus the status of the game is displayed in the console", func(t *testing.T) {
		player1 := &game.PlayerStateMachineImpl{Player: &game.Player{LeftHand: 1, RightHand: 2}}
		player2 := &game.PlayerStateMachineImpl{Player: &game.Player{LeftHand: 3, RightHand: 4}}

		players := []game.PlayerI{player1, player2}
		err := DisplayStatus(players)
		if err != nil {
			t.Errorf("Error: unable to display game status, %v", err)
		}
	})
}
