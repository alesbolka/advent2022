package day13

import (
	"log"
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
	lines := strings.Split(input, "\n")
	pack1 := NewPacket("[[2]]")
	pack2 := NewPacket("[[6]]")
	packets := []*Packet{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		packets = append(packets, NewPacket(line))
	}

	pos1 := 1
	pos2 := 1

	for ii := 0; ii < len(packets); ii++ {
		if pack1.Compare(packets[ii]) == -1 {
			pos1++
		}
		if pack2.Compare(packets[ii]) == -1 {
			pos2++
		}
	}

	switch pack1.Compare(pack2) {
	case 0:
		log.Fatalf("did not expect divider packages to be unsortable")
	case 1:
		pos2++
	case -1:
		pos1++
	}

	return pos1 * pos2
}
