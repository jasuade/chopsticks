package game

import (
	"reflect"
	"strings"
	"testing"
)

func TestPlayTurns(t *testing.T) {
	//Arange
	t.Run("Test PlayTurns when called it sould return an action", func(t *testing.T) {
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

		player := &PlayerOperationsImpl{&Player{1, 1}}
		players := []PlayerI{player, player}

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
	t.Run("Test playSplit that should modify the player correctly", func(t *testing.T) {
		testCase := []struct {
			desc         string
			inputPlayer  *Player
			wantedPlayer *Player
		}{
			{desc: "Test if both are less than or equal 1", inputPlayer: &Player{1, 1}, wantedPlayer: &Player{1, 1}},
			{desc: "Test if player input {4,3} output should be {4,3}", inputPlayer: &Player{4, 3}, wantedPlayer: &Player{4, 3}},
			{desc: "Test if player input {1,2} output should be {2,1}", inputPlayer: &Player{1, 2}, wantedPlayer: &Player{1, 2}},
			{desc: "Test if both are 4", inputPlayer: &Player{4, 4}, wantedPlayer: &Player{4, 4}},
			{desc: "Test if player input {0,2} output should be {1,1}", inputPlayer: &Player{0, 2}, wantedPlayer: &Player{1, 1}},
			{desc: "Test if player input {0,3} output should be {1,2}", inputPlayer: &Player{0, 3}, wantedPlayer: &Player{1, 2}},
			{desc: "Test if player input {1,3} output should return {2,2}", inputPlayer: &Player{1, 3}, wantedPlayer: &Player{2, 2}},
			{desc: "Test if player input {2,2} output should return {3,1}", inputPlayer: &Player{2, 2}, wantedPlayer: &Player{3, 1}},
			{desc: "Test if player input {2,3} output should return {1,4}", inputPlayer: &Player{2, 3}, wantedPlayer: &Player{1, 4}},
			{desc: "Test if player input {1,4} output should return {2,3}", inputPlayer: &Player{1, 4}, wantedPlayer: &Player{2, 3}},
			{desc: "Test if player input {2,4} output should return {3,3}", inputPlayer: &Player{2, 4}, wantedPlayer: &Player{3, 3}},
			{desc: "Test if player input {3,3} output should return", inputPlayer: &Player{3, 3}, wantedPlayer: &Player{4, 2}},
		}
		for _, tcase := range testCase {
			tcase := tcase
			//Act
			t.Run(tcase.desc, func(t *testing.T) {
				t.Parallel()
				playerOperationsImpl := &PlayerOperationsImpl{tcase.inputPlayer}
				playerOperationsImpl.playSplit()
				//Asset
				if !reflect.DeepEqual(tcase.inputPlayer, tcase.wantedPlayer) {
					t.Errorf("Error: got %v but wanted %v\n", tcase.inputPlayer, tcase.wantedPlayer)
				}
			})
		}
	})

	t.Run("Test containsNumber that given a player and a number return true if the player got the number, false otherwise", func(t *testing.T) {
		testCase := []struct {
			desc        string
			inputPlayer *Player
			number      int
			want        bool
		}{
			{desc: "Given a player{0,3} and number 0 return TRUE", inputPlayer: &Player{0, 3}, number: 0, want: true},
			{desc: "Given a player{1,3} and number 0 return FALSE", inputPlayer: &Player{1, 3}, number: 0, want: false},
		}
		for _, tcase := range testCase {
			tcase := tcase
			t.Run(tcase.desc, func(t *testing.T) {
				t.Parallel()
				got := containsNumber(tcase.inputPlayer, tcase.number)
				if got != tcase.want {
					t.Errorf("Error: got %v but wanted %v\n", tcase.inputPlayer, tcase.want)
				}
			})
		}
	})

	t.Run("Test containsFourNumber, given a player{0,4} and input should return the correct split", func(t *testing.T) {
		testCase := []struct {
			desc         string
			inputPlayer  *Player
			inputNumber  int
			wantedPlayer *Player
		}{
			{desc: "Given player {0, 4} and input number 1 should return {1, 3}", inputPlayer: &Player{0, 4}, inputNumber: 1, wantedPlayer: &Player{1, 3}},
			{desc: "Given player {0, 4} and input number 2 should return {2, 2}", inputPlayer: &Player{0, 4}, inputNumber: 2, wantedPlayer: &Player{2, 2}},
			{desc: "Given player {4, 0} and input number 1 should return {3, 1}", inputPlayer: &Player{4, 0}, inputNumber: 1, wantedPlayer: &Player{3, 1}},
			{desc: "Given player {4, 0} and input number 2 should return {2, 2}", inputPlayer: &Player{4, 0}, inputNumber: 2, wantedPlayer: &Player{2, 2}},
		}
		for _, tcase := range testCase {
			tcase := tcase
			//Act
			t.Run(tcase.desc, func(t *testing.T) {
				t.Parallel()
				tcase.inputPlayer.splitWithZeroFour(tcase.inputNumber)
				//Asset
				if !reflect.DeepEqual(tcase.inputPlayer, tcase.wantedPlayer) {
					t.Errorf("Error: got %v but wanted %v\n", tcase.inputPlayer, tcase.wantedPlayer)
				}
			})
		}
	})
}
