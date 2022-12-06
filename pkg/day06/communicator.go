package day06

import (
	"log"
)

func IdentifyStartOfPacket(input string) int {
	return findUniqueSequent(input, 4)
}

func IdentifyStartOfMessage(input string) int {
	return findUniqueSequent(input, 14)
}

func findUniqueSequent(input string, minLength int) int {
	inputLength := len(input)

	if inputLength < minLength {
		log.Fatalf("Input must be at least %d characters long, got %d, encoding issue?", minLength, inputLength)
	}

STRING_LOOP:
	for ii := minLength; ii < inputLength; ii++ {
		dict := map[byte]int{}

		for jj := 0; jj < minLength; jj++ {
			current := input[ii-jj]
			dict[current]++

			if dict[current] > 1 {
				continue STRING_LOOP
			}
		}

		return ii + 1
	}

	// var str segment
	return -1
}
