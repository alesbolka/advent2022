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

func Task2(input string) int {
	sys := newKnotSystem(10)
	var err error
	sys.report.x0 = -11
	sys.report.x1 = 14
	sys.report.y0 = -5
	sys.report.y1 = 15

	// fmt.Println(sys)
	for _, line := range strings.Split(input, "\n") {
		// fmt.Println(line)
		err = sys.ExecuteMove(line)
		// fmt.Println(sys)
		if err != nil {
			log.Fatal(err)
		}
	}

	return len(sys.tailPositions)
}
