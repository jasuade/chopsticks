package game

type Player struct {
	RightHand int
	LeftHand  int
}

func createGame() []Player {
	player1 := Player{1, 1}
	player2 := Player{1, 1}
	return []Player{player1, player2}
}
