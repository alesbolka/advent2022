package day11

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type MonkeyInTheMiddle struct {
	monkeys map[int]*Monkey
	maxId   int
}

func (game *MonkeyInTheMiddle) Round() error {
	var worryValue int
	var target int
	for ii := 0; ii <= game.maxId; ii++ {
		for {
			item, err := game.monkeys[ii].GetItem()
			if err != nil {
				break
			}

			worryValue = game.monkeys[ii].InspectItem(item) / 3
			target = game.monkeys[ii].DetermineTarget(worryValue)
			otherMonkey, exists := game.monkeys[target]
			if !exists {
				return fmt.Errorf("requesting invalid monkey target %d", target)
			}

			// fmt.Printf("Monkey %d throws %d to monkey %d\n", ii, worryValue, target)

			otherMonkey.AddItem(worryValue)
		}
	}

	return nil
}

func (game *MonkeyInTheMiddle) String() (res string) {
	for ii := 0; ii <= game.maxId; ii++ {
		res += fmt.Sprintf("Monkey %d: %s\n", ii, game.monkeys[ii].String())
	}
	return
}

func NewGame(input string) (game *MonkeyInTheMiddle, err error) {
	game = &MonkeyInTheMiddle{
		monkeys: make(map[int]*Monkey),
	}
	lines := strings.Split(input, "\n")
	monkeyIDRegex := regexp.MustCompile(`^Monkey (\d+):$`)
	targetRegex := regexp.MustCompile(`^\s+If (true|false): throw to monkey (\d+)$`)

	var monkey *Monkey
	for _, line := range lines {
		if line == "" {
			continue
		}

		monkeyMatch := monkeyIDRegex.FindStringSubmatch(line)

		if len(monkeyMatch) == 2 {
			id, err := strconv.Atoi(monkeyMatch[1])
			if err != nil {
				return nil, fmt.Errorf("failed parsing monkey id from line '%s'", line)
			}
			monkey = NewMonkey(id)

			if _, exists := game.monkeys[monkey.id]; exists {
				return nil, fmt.Errorf("monkey with the ID %d already exists", monkey.id)
			}

			if monkey.id > game.maxId {
				game.maxId = monkey.id
			}

			game.monkeys[monkey.id] = monkey
			continue
		}

		if strings.HasPrefix(line, "  Starting items: ") {
			lineParts := strings.Split(line, ": ")
			if len(lineParts) != 2 {
				log.Fatalf("Could not parse starting items for '%s'", line)
			}

			for _, item := range strings.Split(lineParts[1], ", ") {
				num, err := strconv.Atoi(item)
				if err != nil {
					log.Fatalf("Failed parsing item number '%s' for line '%s'", item, line)
				}
				monkey.AddItem(num)
			}
			continue
		}

		if strings.HasPrefix(line, "  Operation: ") {
			parts := strings.Split(line, ": new = ")
			if len(parts) != 2 {
				return nil, fmt.Errorf("failed parsing operation line '%s'", line)
			}

			op, err := ParseOperation(parts[1])
			if err != nil {
				return nil, err
			}
			monkey.logic = op
			continue
		}

		if strings.HasPrefix(line, "  Test: divisible by ") {
			parts := strings.Split(line, "Test: divisible by ")
			if len(parts) != 2 {
				return nil, fmt.Errorf("failed parsing test line '%s'", line)
			}

			val, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("failed determining divisor '%s'", line)
			}
			monkey.divisor = val
			continue
		}

		targetMatch := targetRegex.FindStringSubmatch(line)
		if len(targetMatch) == 3 {
			val, err := strconv.Atoi(targetMatch[2])
			if err != nil {
				return nil, fmt.Errorf("failed parsing target from '%s'", line)
			}
			switch targetMatch[1] {
			case "false":
				monkey.targets[0] = val
			case "true":
				monkey.targets[1] = val
			}
			continue
		}

		log.Fatalf("Failed to process line '%s'", line)
	}

	for ii := 0; ii <= game.maxId; ii++ {
		if _, exists := game.monkeys[ii]; !exists {
			return nil, fmt.Errorf("monkey parsing out of order, cannot find monkey %d", ii)
		}
	}
	return
}
