package day15

import (
	"log"
)

func Task1(input string) int {
	arr, err := NewArray(input)
	if err != nil {
		log.Fatal(err)
	}
	return arr.CountEmptySpots(2000000)
}

func Task2(input string) int64 {
	arr, err := NewArray(input)
	if err != nil {
		log.Fatal(err)
	}

	solutions := arr.FindEmptySpot(0, 4000000)
	if len(solutions) != 1 {
		log.Fatalf("Did not find 1 solution: %v", solutions)
	}

	for key := range solutions {
		return int64(key.x)*4000000 + int64(key.y)
	}
	return -1
}
