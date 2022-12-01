package main

import (
	"fmt"

	asset "github.com/alesbolka/advent2022/assets/day01"
	"github.com/alesbolka/advent2022/pkg/day01"
)

func main() {
	caloryElf, err := day01.FindHighCalorieElf(asset.Input)
	if err != nil {
		fmt.Println("Failed finding the elf with most calories", err)
		return
	}
	fmt.Printf("One elf is carrying %d calories\n", caloryElf)

	top3Calories, err := day01.FindTopThreeCaloricElves(asset.Input)

	if err != nil {
		fmt.Println("Failed finding the top 3 elves with most calories", err)
		return
	}

	fmt.Printf("The 3 elves with most snacks are carrying %d calories\n", top3Calories)
}
