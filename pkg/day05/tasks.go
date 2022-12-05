package day05

import (
	"log"
	"strings"
)

func Task1(input string) string {
	crateState := []string{}
	// instructions := []string{}
	handlingState := true
	var stk *stockpile

	for ii, line := range strings.Split(input, "\n") {
		if strings.TrimSpace(line) == "" {
			if !handlingState {
				log.Fatalf("Found empty second empty line on line %d, aborting", ii)
			}

			handlingState = false
			stk = parseStockpile(crateState[len(crateState)-1], crateState[:len(crateState)-1])
			continue

		}

		if handlingState {
			crateState = append(crateState, line)
			continue
		}

		// log.Printf("before %s\n%v\n", line, stk)
		stk.ExecuteCommand(line)
		// log.Printf("after\n%v\n", stk)
	}

	return stk.GetTop()
}

func Task2(input string) (res int) {
	return
}
