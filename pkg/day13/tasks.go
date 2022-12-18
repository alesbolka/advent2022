package day13

import (
	"strings"
)

func Task1(input string) int {
	lines := strings.Split(input, "\n")
	packets := []*Packet{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		packets = append(packets, NewPacket(line))
	}

	res := 0
	for ii := 0; ii < len(packets); ii += 2 {
		if packets[ii].Compare(packets[ii+1]) == 1 {
			// fmt.Println(ii + 1)
			res += 1 + (ii / 2)
		}
	}
	return res
}

func Task2(input string) int {
	return 0
}
