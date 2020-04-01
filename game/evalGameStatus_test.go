package game

import (
	"testing"

	game "github.hc.ag/jsuarez/chopsticks/game/operationsImplementation"
)

func TestEvalGameStatus(t *testing.T) {
	t.Run("", func(t *testing.T) {
		player := &game.PlayerOperationsImpl{&game.Player{1, 1}}
		players := []game.PlayerI{player, player}
		got := isWin(players)

		want := true

		if want != got {
			t.Errorf("Error: got %v but wanted %v\n", got, want)
		}
	})
}
