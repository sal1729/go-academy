package utils

import "testing"

type RollTestCandidate struct {
	RollSum         int
	ExpectedOutcome string
}

var RollTestCandidates = []RollTestCandidate{
	{2, "SNAKE-EYES-CRAPS"},
	{3, "LOSS-CRAPS"},
	{4, "NEUTRAL"},
	{5, "NEUTRAL"},
	{6, "NEUTRAL"},
	{7, "NATURAL"},
	{8, "NEUTRAL"},
	{9, "NEUTRAL"},
	{10, "NEUTRAL"},
	{11, "NATURAL"},
	{12, "LOSS-CRAPS"},
}

func TestAllPossibleRolls(t *testing.T) {
	for _, candidate := range RollTestCandidates {
		outcome := rollOutcome(candidate.RollSum)
		if outcome != candidate.ExpectedOutcome {
			t.Errorf("For roll %d expected %s, got %s", candidate.RollSum, candidate.ExpectedOutcome, outcome)
		}
	}
}
