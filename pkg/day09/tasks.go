package day09

import (
	"log"
	"strings"
)

func Task1(input string) uint64 {
	sys := NewSystem(0, 0)
	var err error
	for _, line := range strings.Split(input, "\n") {
		err = sys.ExecuteMove(line)
		if err != nil {
			log.Fatal(err)
		}
	}
	return sys.VisitedCounter
}
