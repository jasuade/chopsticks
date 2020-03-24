package ui

import (
	"testing"

	"github.hc.ag/jsuarez/chopsticks/game"
)

func TestUIDisplaying(t *testing.T) {
	t.Run("When call DisplayStatus the status of the game is displayed in the console", func(t *testing.T) {
		players := []game.Player{game.Player{5, 3}, game.Player{1, 4}}
		err := DisplayStatus(players)
		if err != nil {
			t.Errorf("Error: unable to display game status, %v", err)
		}
	})
}
