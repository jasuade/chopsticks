package game

import (
	"strings"
	"testing"
)

func TestPlayTurns(t *testing.T) {
	//Arange
	t.Run("When called the function PlayTurns it sould return an action", func(t *testing.T) {
		testCase := []struct {
			desc  string
			input string
			want  string
		}{
			{desc: "Selected the action a we should recieve 'attack'", input: "a", want: "attack"},
			{desc: "Selected the action s we should recieve 'split'", input: "s", want: "split"},
			{desc: "Selected the any other letter or num we should recieve 'Invalid action'", input: "3", want: "Invalid action"},
			{desc: "Selected the any other letter or num we should recieve 'Invalid action'", input: "sd", want: "Invalid action"},
		}
		players := []Player{Player{1, 1}, Player{1, 1}}

		for _, tcase := range testCase {
			tcase := tcase
			//Act
			t.Run(tcase.desc, func(t *testing.T) {
				t.Parallel()
				reader := strings.NewReader(tcase.input)
				got := playTurnFromReader(players, 1, reader)
				//Asset
				if got != tcase.want {
					t.Errorf("Error: got %v but wanted %v\n", got, tcase.want)
				}
			})
		}
	})
}
