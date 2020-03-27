package game

import "fmt"

type PlayerStateMachineImpl struct {
	player *Player
}

func (psi *PlayerStateMachineImpl) GetPlayer() *Player {
	return psi.player
}

func (psi *PlayerStateMachineImpl) playAttack() {
}

func (psi *PlayerStateMachineImpl) playSplit() {
	fmt.Println("Mierda pura")
}
