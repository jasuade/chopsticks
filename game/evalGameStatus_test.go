package game

import (
	"testing"

	game "github.hc.ag/jsuarez/chopsticks/game/operationsImplementation"
)

func TestEvalGameStatus(t *testing.T) {
	t.Run("Should return false when calling IsWin with no winner player", func(t *testing.T) {
		player := &game.PlayerOperationsImpl{&game.Player{1, 1}}
		players := []game.PlayerI{player, player}
		got := IsWin(players)

		want := false

		if want != got {
			t.Errorf("Error: got %v but wanted %v\n", got, want)
		}
	})

	t.Run("Should return true when calling IsWin with a winner player", func(t *testing.T) {
		player1 := &game.PlayerOperationsImpl{&game.Player{1, 1}}
		player2 := &game.PlayerOperationsImpl{&game.Player{0, 0}}
		players := []game.PlayerI{player1, player2}
		got := IsWin(players)

		want := true

		if want != got {
			t.Errorf("Error: got %v but wanted %v\n", got, want)
		}
	})
}
