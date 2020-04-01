package game

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

//PlayTurn execute the action of a player in a given moment, it uses the stdio to read from the console
func PlayTurn(players []PlayerI, playerTurn int) error {
	return playTurnFromReader(players, playerTurn, os.Stdin)
}

func playTurnFromReader(players []PlayerI, playerTurn int, r io.Reader) error {
	reader := bufio.NewReader(r)

	fmt.Print("Are you goint to attack(a) or to split(s):\n")
	action, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	switch strings.TrimSpace(action) {
	case "a":
		fmt.Println("You choose to attack")
		err = chooseAttack(players, playerTurn, os.Stdin)
	case "s":
		fmt.Println("You choose to split")
		err = players[playerTurn].playSplit()
	default:
		err := errors.New("Invalid action: please introduce a valid action")
		return err
	}
	return err
}
