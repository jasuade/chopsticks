package game

import (
	"strings"
	"testing"
)

func TestPlayTurns(t *testing.T) {
	//Arange
	t.Run("Test PlayTurns when called it sould perform the action", func(t *testing.T) {
		testCase := []struct {
			desc  string
			input string
		}{
			//{desc: "Selected the action a should perform an attack", input: "a\nl\n"}, //Mock choose Attack to return nil
			{desc: "Selected the action s should perfomr a split", input: "s\n"},
		}

		player := &PlayerOperationsImpl{&Player{2, 2}}
		players := []PlayerI{player, player}

		for _, tcase := range testCase {
			tcase := tcase
			//Act
			t.Run(tcase.desc, func(t *testing.T) {
				t.Parallel()
				reader := strings.NewReader(tcase.input)
				err := playTurnFromReader(players, 1, reader)
				//Asset
				if err != nil {
					t.Errorf("Got an error %v when reding", err)
				}
			})
		}
	})
}
