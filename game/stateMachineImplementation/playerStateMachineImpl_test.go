package game

import (
	"reflect"
	"testing"
)

func TestPlaySplitStateMachineImpl(t *testing.T) {
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
				playerOperationsImpl := &PlayerStateMachineImpl{tcase.inputPlayer}
				playerOperationsImpl.playSplit()
				//Asset
				if !reflect.DeepEqual(tcase.inputPlayer, tcase.wantedPlayer) {
					t.Errorf("Error: got %v but wanted %v\n", tcase.inputPlayer, tcase.wantedPlayer)
				}
			})
		}
	})
}
