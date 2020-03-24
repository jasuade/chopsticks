package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//PlayTurn executed the action of a player in a given moment, it uses the stdio to read from the console
func PlayTurn(players []Player, playerTurn int) string {
	return playTurnFromReader(players, playerTurn, os.Stdin)
}

func playTurnFromReader(players []Player, playerTurn int, r io.Reader) string {
	reader := bufio.NewReader(r)

	fmt.Print("Are you goint to attack(a) or to split(s):\n")

	action, _ := reader.ReadString('\n')

	// if err != nil {
	// 	return "err"
	// }

	switch strings.TrimSpace(action) {
	case "a":
		fmt.Println("You choose to attack")
		action = "attack"
		//attack()
	case "s":
		fmt.Println("You choose to split")
		action = "split"
		//split()
	default:
		action = "Invalid action"
		fmt.Println(action)
	}
	return action
}
