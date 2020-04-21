package ui

import game "github.hc.ag/jsuarez/chopsticks/game/operationsImplementation"

const ATTACK_OR_SPLIT byte = 0
const LEFT_OR_RIGHT_ATTACKER byte = 1
const LEFT_OR_RIGHT_DEFENDER byte = 2
const ONE_OR_TWO byte = 3
const IS_WIN byte = 4

type GameState struct {
	Players     []game.PlayerI
	PlayerTurn  int
	Msg         string
	AttackPower int
	TurnState   byte
}

type MessagesPrinter interface {
	PrintMessage(text string) error
}

var Messages = map[string]string{
	"MSG_START":               "Start the game with %d payers in mode <%s> and level <%s>",
	"MSG_TURN":                "Player %d is your turn",
	"MSG_HAND_TO_ATTACK_WITH": "With which hand (left(l) or right(r)) do you want to attack:",
	"MSG_HAND_TO_ATTACK_TO":   "And which hand (left(l) or right(r)) do you want to attack:",
	"MSG_SPLIT":               "How many chopsticks you want to transfere?",
	"MSG_ACTION":              "Are you goint to attack(a) or to split(s):",
	"MSG_WIN":                 "Player P%d WINS the game!!!",
	"MSG_LOSE":                "Player P%d LOSES the game",
	"ERR_INVALID_ACTION":      "That is not a valid action, please chose attack (a) or split (s)",
	"ERR_INVALID_ATTACK":      "Invalid attack, the hand does not exist",
	"ERR_SPLIT_HAND":          "Unable to slpit, you cannot kill a hand",
	"ERR_SPLIT":               "That is not a meaninful split",
	"ERR_ACTION":              "Invalid action: please introduce a valid action",
}
