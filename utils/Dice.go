package utils

import (
	"math/rand"
)

// TODO Write a multi-case test
func rollOutcome(roll [2]int) string {
	// switch on the roll values
	switch roll[0] + roll[1] {
	case 2:
		return "SNAKE-EYES-CRAPS"
	case 3, 12:
		return "LOSS-CRAPS"
	case 7, 11:
		return "NATURAL"
	default:
		return "NEUTRAL"
	}
}

func RollDicePair() ([2]int, string) {
	roll := [2]int{rand.Intn(6) + 1, rand.Intn(6) + 1}
	return roll, rollOutcome(roll)
}
