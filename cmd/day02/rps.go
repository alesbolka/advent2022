package main

import (
	"fmt"

	asset "github.com/alesbolka/advent2022/assets/day02"
	"github.com/alesbolka/advent2022/pkg/day02"
)

func main() {
	res, err := day02.CalculateStrategyScore(asset.Input)

	fmt.Println(res, err)
}
