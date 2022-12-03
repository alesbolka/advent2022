package day03

import (
	"strings"
)

func Task1(input string) (sum int) {
	for _, line := range strings.Split(input, "\n") {
		bkp := newBackpack(line)
		sum += bkp.getDuplicatePriority()
	}
	return
}

func Task2(input string) (sum int) {
	var group elfGroup
	var elfIndex int

	for ii, line := range strings.Split(input, "\n") {
		elfIndex = ii % 3

		if elfIndex == 0 {
			group = elfGroup{[3]*backpack{}}
		}

		group.packs[elfIndex] = newBackpack(line)
		if elfIndex == 2 {
			sum += group.findBadgePriority()
		}
	}
	return
}
