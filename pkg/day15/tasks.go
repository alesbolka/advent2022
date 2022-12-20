package day15

import "log"

func Task1(input string, targetRow int) int {
	arr, err := NewArray(input)
	if err != nil {
		log.Fatal(err)
	}
	return arr.CountEmptySpots(targetRow)
}

func Task2(input string) int {
	return 0
}
