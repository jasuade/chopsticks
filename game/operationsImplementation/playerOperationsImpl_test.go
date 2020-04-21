package game

import (
	"reflect"
	"strings"
	"testing"
)

func TestPlaySplitOperationsImpl(t *testing.T) {
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
				playerOperationsImpl.PlaySplit()
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
				got := ContainsNumber(tcase.inputPlayer, tcase.number)
				if got != tcase.want {
					t.Errorf("Error: got %v but wanted %v\n", tcase.inputPlayer, tcase.want)
				}
			})
		}
	})

}

func TestPlayAttacktOperationsImpl(t *testing.T) {
	t.Run("Test play Attack, given a player, the attacker hand and the oponent hand return a oponent player", func(t *testing.T) {
		// chooseAttack(players []PlayerI, playerTurn int, r io.Reader)
		// (poi *PlayerOperationsImpl) playAttack(defendingPlayer *Player, num int, receiverHand string) {
		testCase := []struct {
			desc                string
			attackerPlayer      *Player
			oponentPlayerInput  *Player
			oponentPlayerOutput *Player
			attackerHand        string
			oponentHand         string
		}{
			{desc: "Given an attacker player{0,3} and a oponent{3,2}, an attack hand r, and a oponent hand r, return oponent{3,5}", attackerPlayer: &Player{0, 3}, oponentPlayerInput: &Player{3, 2},
				attackerHand: "r", oponentHand: "r", oponentPlayerOutput: &Player{3, 0}},
			{desc: "Given an attacker player{2,1} and a oponent{2,2}, an attack hand l, and a oponent hand r, return oponent{2, 4}", attackerPlayer: &Player{2, 1}, oponentPlayerInput: &Player{2, 2},
				attackerHand: "l", oponentHand: "r", oponentPlayerOutput: &Player{2, 4}},
			{desc: "Given an attacker player{0,3} and a oponent{3,2}, an attack hand l, and a oponent hand r, return the same oponent", attackerPlayer: &Player{0, 3}, oponentPlayerInput: &Player{3, 2},
				attackerHand: "l", oponentHand: "r", oponentPlayerOutput: &Player{3, 2}},
			{desc: "Given an attacker player{1,2} and a oponent{4,0}, an attack hand r, and a oponent hand r, return the same oponent", attackerPlayer: &Player{1, 2}, oponentPlayerInput: &Player{4, 0},
				attackerHand: "r", oponentHand: "r", oponentPlayerOutput: &Player{4, 0}},
		}
		for _, tcase := range testCase {
			tcase := tcase
			t.Run(tcase.desc, func(t *testing.T) {
				t.Parallel()

				hands := strings.NewReader(tcase.attackerHand + "\n" + tcase.oponentHand + "\n")

				p0 := &PlayerOperationsImpl{tcase.attackerPlayer}
				p1 := &PlayerOperationsImpl{tcase.oponentPlayerInput}
				players := []PlayerI{p0, p1}

				turn := 0
				chooseAttack(players, turn, hands)
				if !reflect.DeepEqual(tcase.oponentPlayerInput, tcase.oponentPlayerOutput) {
					t.Errorf("Error: got %v but wanted %v\n", tcase.oponentPlayerInput, tcase.oponentPlayerOutput)
				}
			})
		}
	})

}
