package game

import (
	"fmt"
	"strconv"
	"strings"
)

//PlayerStateMachineImpl is the current implementation of the interface PlayerI, implementation done with math operations
type PlayerStateMachineImpl struct {
	Player *Player
}

var PlayerStates = map[string]string{
	"22": "31",
	"13": "22",
	"23": "14",
	"14": "23",
	"33": "42",
	"24": "33",
	"02": "11",
	"03": "12",
	"04": "22",
	//"04": "13",
}

//GetPlayer function restuns the Current implementation of the Player
func (psi *PlayerStateMachineImpl) GetPlayer() *Player {
	return psi.Player
}

func (psi *PlayerStateMachineImpl) playAttack() {
}

func (psi *PlayerStateMachineImpl) playSplit() {
	player := psi.Player
	state := strconv.Itoa(player.LeftHand) + strconv.Itoa(player.RightHand)
	if value, ok := PlayerStates[state]; ok {
		tmp := strings.Split(value, "")

		player.LeftHand, _ = strconv.Atoi(tmp[0])
		player.RightHand, _ = strconv.Atoi(tmp[1])
		return

	}
	fmt.Println("Invalid split")
}
