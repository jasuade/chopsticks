package game

import (
	"reflect"
	"testing"

	"github.hc.ag/jsuarez/chopsticks/config"
)

func TestReadConfiguration(t *testing.T) {
	t.Run("When call createGame should return 2 players with 1 finger each", func(t *testing.T) {
		config := &config.Configuration{2, "easy", "cutoff"}
		got := createGame(config)
		want := []Player{Player{1, 1}, Player{1, 1}}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Error: got %v but wanted %v\n", got, want)
		}
	})
}
