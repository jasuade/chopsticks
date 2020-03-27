package game

import (
	"reflect"
	"testing"

	"github.hc.ag/jsuarez/chopsticks/config"
)

func TestCreateGame(t *testing.T) {
	t.Run("When call createGame should return 2 players with 1 finger each", func(t *testing.T) {
		config := &config.Configuration{2, "easy", "cutoff"}
		got := CreateGame(config)
		player := &PlayerOperationsImpl{&Player{1, 1}}
		want := []PlayerI{player, player}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Error: got %v but wanted %v\n", got, want)
		}
	})

	t.Run("When call startGame should return one random players to start", func(t *testing.T) {
		config := &config.Configuration{2, "easy", "cutoff"}
		got := StartGame(config)
		want := []int{0, 1}
		if got != want[0] && got != want[1] {
			t.Errorf("Error: got %v but wanted %v\n", got, want)
		}
	})
}
