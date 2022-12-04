package day04

import (
	"log"
	"strconv"
	"strings"
)

type assignment struct {
	areas [2]struct {
		min int
		max int
	}
}

func interpretInstruction(line string) *assignment {
	pair := strings.Split(line, ",")
	if len(pair) != 2 {
		log.Fatalf("Line %s did not parse", line)
	}

	res := &assignment{}
	var err error

	for ii := 0; ii < 2; ii++ {
		coords := strings.Split(pair[ii], "-")
		if len(coords) != 2 {
			log.Fatalf("Line %s did not have correct pair at %d", line, ii)
		}

		res.areas[ii].min, err = strconv.Atoi(coords[0])
		if err != nil {
			log.Fatalf("Failed integer parse for %s; %v", line, err)
		}

		res.areas[ii].max, err = strconv.Atoi(coords[1])
		if err != nil {
			log.Fatalf("Failed integer parse for %s; %v", line, err)
		}
	}

	return res
}

func (asg *assignment) isRangeFullyContained() bool {

	res := asg.areas[0].min >= asg.areas[1].min && asg.areas[0].max <= asg.areas[1].max ||
		asg.areas[1].min >= asg.areas[0].min && asg.areas[1].max <= asg.areas[0].max

	return res
}
