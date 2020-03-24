package game

import (
	"reflect"
	"testing"

	"github.hc.ag/jsuarez/chopsticks/config"
)

func TestReadConfiguration(t *testing.T) {
	t.Run("When call createGame should return 2 players with 1 finger each", func(t *testing.T) {
		config := &config.Configuration{2, "easy", "cutoff"}
		got := CreateGame(config)
		want := []Player{Player{1, 1}, Player{1, 1}}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Error: got %v but wanted %v\n", got, want)
		}
	})

	t.Run("When call startGame should return one random players to start", func(t *testing.T) {
		config := &config.Configuration{2, "easy", "cutoff"}
		players := []Player{Player{1, 1}, Player{1, 1}}
		got := StartGame(players, config)
		want := []int{0, 1}
		if got != want[0] && got != want[1] {
			t.Errorf("Error: got %v but wanted %v\n", got, want)
		}
	})
}
