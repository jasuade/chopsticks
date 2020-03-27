package game

//PlayerStateMachineImpl is the current implementation of the interface PlayerI, implementation done with math operations
type PlayerStateMachineImpl struct {
	player *Player
}

//GetPlayer function restuns the Current implementation of the Player
func (psi *PlayerStateMachineImpl) GetPlayer() *Player {
	return psi.player
}

func (psi *PlayerStateMachineImpl) playAttack() {
}

func (psi *PlayerStateMachineImpl) playSplit() {
}
