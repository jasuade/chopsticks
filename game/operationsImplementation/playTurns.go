package game

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

//PlayTurn execute the action of a player in a given moment, it uses the stdio to read from the console
func PlayTurn(players []PlayerI, playerTurn int) error {
	return playTurnFromReader(players, playerTurn, os.Stdin)
}

func playTurnFromReader(players []PlayerI, playerTurn int, r io.Reader) error {
	reader := bufio.NewReader(r)

	log.Print("Are you goint to attack(a) or to split(s):\n")
	action, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	switch strings.TrimSpace(action) {
	case "a":
		err = chooseAttack(players, playerTurn, r)
	case "s":
		err = players[playerTurn].PlaySplit()
	default:
		err := errors.New("ERR_ACTION")
		return err
	}
	return err
}
