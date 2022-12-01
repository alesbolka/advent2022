package day01

import (
	"strconv"
	"strings"
)

func FindHighCalorieElf(input string) (max int, err error) {
	max = 0
	err = nil

	lines := strings.Split(input, "\n")

	elf := 0

	for ii, line := range lines {
		if len(line) < 1 {
			if elf > max {
				max = elf
			}
			elf = 0
			continue
		}

		mealCalories, err := strconv.Atoi(line)
		if err != nil {
			return ii, err
		}

		elf += mealCalories
	}

	if elf > max {
		max = elf
	}

	return
}

func FindTopThreeCaloricElves(input string) (int, error) {
	lines := strings.Split(input, "\n")
	top3 := [3]int{0, 0, 0}
	elf := 0

	rankElf := func(elfCalories int) {
		for ii := 0; ii < 3; ii++ {
			if elf > top3[ii] {
				for jj := 1; jj >= ii; jj-- {
					top3[jj+1] = top3[jj]
				}
				top3[ii] = elf
				break
			}
		}
	}

	for ii, line := range lines {
		if len(line) < 1 {
			rankElf(elf)
			elf = 0
			continue
		}

		mealCalories, err := strconv.Atoi(line)
		if err != nil {
			return ii, err
		}

		elf += mealCalories
	}

	rankElf(elf)

	return top3[0] + top3[1] + top3[2], nil
}
