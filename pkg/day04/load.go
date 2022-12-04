package day04

import "strings"

func Task1(input string) (res int) {
	for _, line := range strings.Split(input, "\n") {
		if interpretInstruction(line).isRangeFullyContained() {
			res++
		}
	}

	return
}

func Task2(input string) (res int) {
	for _, line := range strings.Split(input, "\n") {
		if interpretInstruction(line).doesRangeOverlap() {
			res++
		}
	}

	return
}
