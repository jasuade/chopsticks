package config

import (
	"os"
	"reflect"
	"testing"
)

func TestReadConfiguration(t *testing.T) {
	t.Run("Test to open the file", func(t *testing.T) {
		_, err := OpenConfig("game.config.test")
		if err != nil {
			t.Errorf("Error %v when opening the file", err)
		}
	})

	t.Run("Test to read the first line of the file", func(t *testing.T) {
		want := &Configuration{
			2,
			"easy",
			"cutoff",
		}
		var file, err = os.OpenFile("game.config.test", os.O_RDONLY, 0644)
		if err != nil {
			t.Errorf("Error %v when opening the file", err)
		}
		got := ReadConfig(file)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Error: got %v but wanted %v\n", got, want)
		}
	})
}
