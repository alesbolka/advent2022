package day05

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// Warning: this only works with stockpiles with less than 10 stacks
type stockpile struct {
	stacks   map[byte][]byte
	baseLine []byte
}

var instructionRegex regexp.Regexp = *regexp.MustCompile(`^move (\d+) from (\d)+ to (\d)+$`)

func parseStockpile(stacks string, distribution []string) (pile *stockpile) {

	maxPos := 0

	positionMap := map[byte]int{}

	stackLength := len(stacks)
	for ii := 0; ii < stackLength; ii++ {
		if stacks[ii] != ' ' {
			if (ii+1) < stackLength && stacks[ii+1] != ' ' {
				log.Fatal("this system can only parse single character position markers")
			}

			positionMap[stacks[ii]] = ii
			maxPos = ii
		}
	}

	pile = &stockpile{
		stacks:   map[byte][]byte{},
		baseLine: make([]byte, maxPos*2),
	}

	for ii := 0; ii < maxPos*2; ii++ {
		pile.baseLine[ii] = ' '
	}

	for ii := len(distribution) - 1; ii >= 0; ii-- {
		maxIndex := len(distribution[ii]) - 1

		for designationChar, positionInt := range positionMap {
			if positionInt > maxIndex {
				// the string has no content for this index, saving input with vscode trims empty spaces
				continue
			}

			if distribution[ii][positionInt] == ' ' {
				continue
			}

			pile.stacks[designationChar] = append(pile.stacks[designationChar], distribution[ii][positionInt])
		}

	}

	return
}

func (pile *stockpile) ExecuteCommand(command string, model int) {
	matched := instructionRegex.FindStringSubmatch(command)

	if len(matched) != 4 {
		log.Fatalf("Could not parse instruction \"%s\" correctly", command)
	}

	count, _ := strconv.Atoi(matched[1])
	from := matched[2][0]
	to := matched[3][0]

	var moving byte

	if model == 9000 {
		// Very unsafe, no checks for length
		for ii := 0; ii < count; ii++ {
			indexOfLast := len(pile.stacks[from]) - 1

			moving, pile.stacks[from] = pile.stacks[from][indexOfLast], pile.stacks[from][:indexOfLast]
			pile.stacks[to] = append(pile.stacks[to], moving)
		}
		return
	}

	if model != 9001 {
		log.Fatalf("Unsupported model %d", model)
	}

	// Very unsafe, no checks for length
	fromlength := len(pile.stacks[from])
	pile.stacks[to] = append(pile.stacks[to], pile.stacks[from][fromlength-count:]...)
	pile.stacks[from] = pile.stacks[from][:fromlength-count]
}

func (pile *stockpile) GetTop() (res string) {
	for ii := 0; ii < len(pile.stacks); ii++ {
		index := '1' + byte(ii)
		if len(pile.stacks[index]) == 0 {
			fmt.Println("empty stack")
			continue
		}
		res += string(pile.stacks[index][len(pile.stacks[index])-1])
	}
	return
}

func (pile *stockpile) String() string {
	maxL := 0
	for _, stack := range pile.stacks {
		if len(stack) > maxL {
			maxL = len(stack)
		}
	}

	out := ""

	for ii := maxL - 1; ii >= 0; ii-- {
		line := make([]byte, len(pile.baseLine))
		copy(line, pile.baseLine)

		for pos, stack := range pile.stacks {
			if ii >= len(stack) {
				continue
			}
			index, _ := strconv.Atoi(string(pos))

			line[(index-1)*2] = stack[ii]
		}

		out += string(line) + "\n"
	}

	for ii := 1; ii <= len(pile.stacks); ii++ {
		out += fmt.Sprintf("%d ", ii)
	}

	return out
}
