package day02

import (
	"fmt"
	"strings"
)

type rpsValue struct {
	initialValue int
	myValue      map[byte]int
}

var scoreMap map[byte]map[byte]int = map[byte]map[byte]int{
	// Rock
	'A': {
		'X': 3, // we tie
		'Y': 6, // my paper wins
		'Z': 0, // my scissors lose
	},
	// Paper
	'B': {
		'X': 0, // my rock loses
		'Y': 3, // we tie
		'Z': 6, // my scissors win
	},
	// Scissors
	'C': {
		'X': 6, // my rock wins
		'Y': 0, // my paper loses
		'Z': 3, // we tie
	},
}

func CalculateStrategyScore(input string) (int, error) {
	lines := strings.Split(input, "\n")
	score := 0

	var lineScore int
	var err error

	for _, line := range lines {
		if len(line) != 3 {
			// fmt.Println("skipping line")
			continue
		}

		lineScore, err = calcScore(line[0], line[2])
		if err != nil {
			return -1, err
		}

		score += lineScore
	}
	return score, nil
}

func calcScore(opponent, myChoice byte) (int, error) {
	firstMatch, exists := scoreMap[opponent]

	if !exists {
		return -1, fmt.Errorf("invalid symbol %s", string(opponent))
	}

	winScore, exists := firstMatch[myChoice]
	if !exists {
		return -1, fmt.Errorf("invalid symbol %s", string(myChoice))
	}

	switch myChoice {
	case 'X':
		return 1 + winScore, nil
	case 'Y':
		return 2 + winScore, nil
	case 'Z':
		return 3 + winScore, nil
	}

	return 0, fmt.Errorf("should never happen")
}

func CalculatePredictionScore(input string) (int, error) {
	lines := strings.Split(input, "\n")
	var lineScore int
	var err error
	score := 0

	for _, line := range lines {
		if len(line) != 3 {
			continue
		}

		choice := 'a'
		switch line[2] {
		case 'X': // need to lose
			switch line[0] {
			case 'A':
				choice = 'Z' // pick scissors against rock
			case 'B':
				choice = 'X' // pick rock against paper
			case 'C':
				choice = 'Y' // pick paper against scissors
			}
		case 'Y': // need to draw
			switch line[0] {
			case 'A':
				choice = 'X'
			case 'B':
				choice = 'Y'
			case 'C':
				choice = 'Z'
			}
		case 'Z': // need to win
			switch line[0] {
			case 'A':
				choice = 'Y'
			case 'B':
				choice = 'Z'
			case 'C':
				choice = 'X'
			}
		default:
			return -1, fmt.Errorf("invalid value for match prediction %s", string(line[2]))
		}

		lineScore, err = calcScore(line[0], byte(choice))
		if err != nil {
			return -1, err
		}
		score += lineScore
	}

	return score, nil
}
