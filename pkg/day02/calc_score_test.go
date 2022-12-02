package day02

import (
	"testing"
)

var opponentRock rune = 'A'
var myRock rune = 'X'
var opponentPaper rune = 'B'
var myPaper rune = 'Y'
var opponentScissors rune = 'C'
var myScissors rune = 'Z'

func Test_scoreCalculator(test *testing.T) {
	type seed struct {
		opponent rune
		me       rune
		expected int
	}
	data := []seed{
		{opponentRock, myRock, 1 + 3},
		{opponentRock, myPaper, 2 + 6},
		{opponentRock, myScissors, 3 + 0},
		{opponentPaper, myRock, 1 + 0},
		{opponentPaper, myPaper, 2 + 3},
		{opponentPaper, myScissors, 3 + 6},
		{opponentScissors, myRock, 1 + 6},
		{opponentScissors, myPaper, 2 + 0},
		{opponentScissors, myScissors, 3 + 3},
	}

	for _, row := range data {
		score, err := calcScore(byte(row.opponent), byte(row.me))
		if err != nil {
			test.Fatalf("%s vs %s, got error %v", string(row.opponent), string(row.me), err)
		}

		if score != row.expected {
			test.Fatalf(
				"%s vs %s, expected %d got error %d",
				string(row.opponent),
				string(row.me),
				row.expected,
				score,
			)
		}
	}
}
