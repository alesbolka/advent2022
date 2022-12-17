package day11

import (
	"errors"
	"fmt"
	"strings"
)

type Monkey struct {
	divisor      int
	id           int
	inspectCount int
	items        []int
	logic        *operation
	targets      [2]int
}

func NewMonkey(id int) *Monkey {
	return &Monkey{
		id:      id,
		items:   []int{},
		targets: [2]int{},
	}
}

func (monkey *Monkey) AddItem(item int) {
	monkey.items = append(monkey.items, item)
}

func (monkey *Monkey) GetItem() (int, error) {
	if len(monkey.items) == 0 {
		return 0, errors.New("no items")
	}
	var item int
	item, monkey.items = monkey.items[0], monkey.items[1:]

	return item, nil
}

func (monkey *Monkey) InspectItem(item int) int {
	monkey.inspectCount++
	return monkey.logic.operand(
		monkey.logic.arg1(item),
		monkey.logic.arg2(item),
	)
}

func (monkey *Monkey) DetermineTarget(item int) int {
	// fmt.Printf("Monkey %d dividing %d by %d, got remainder %d\n", monkey.id, item, monkey.divisor, item%monkey.divisor)
	if (item % monkey.divisor) == 0 {
		return monkey.targets[1]
	}
	return monkey.targets[0]
}

func (monkey *Monkey) String() string {
	itemL := len(monkey.items)
	strs := make([]string, itemL)
	for ii := 0; ii < itemL; ii++ {
		strs[ii] = fmt.Sprintf("%d", monkey.items[ii])
	}

	return strings.Join(strs, ", ")
}
