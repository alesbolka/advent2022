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
