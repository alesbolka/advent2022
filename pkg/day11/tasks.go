package day11

import (
	"log"
)

func Task1(input string) int {
	game, err := NewGame(input)
	if err != nil {
		log.Fatal(err)
	}
	for ii := 0; ii < 20; ii++ {
		game.Round()
		// fmt.Println(game)
	}

	max := [2]int{0, 0}
	for _, monkey := range game.monkeys {
		if monkey.inspectCount > max[1] {
			max[0], max[1] = max[1], monkey.inspectCount
			continue
		}

		if monkey.inspectCount > max[0] {
			max[0] = monkey.inspectCount
		}
	}

	return max[0] * max[1]
}

func Task2(input string) string {
	return ""
}
