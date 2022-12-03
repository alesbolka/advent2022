package day03

import (
	"fmt"
	"log"
)

type backpack struct {
	original     string
	compartments [2]map[byte]int
	duplicates   map[byte]int
	global       map[byte]int
}

func newBackpack(contents string) (bkp *backpack) {
	contentLength := len(contents)
	if contentLength%2 > 0 {
		log.Fatalf("uneven distribution for backpack %s", contents)
	}

	bkp = &backpack{
		original:     contents,
		compartments: [2]map[byte]int{{}, {}},
		duplicates:   map[byte]int{},
		global:       map[byte]int{},
	}

	for ii := 0; ii < contentLength; ii++ {
		item := contents[ii]
		index := 0
		if ii >= contentLength/2 {
			index = 1

			if bkp.compartments[0][item] > 0 {
				bkp.duplicates[item]++
			}
		}

		bkp.global[item]++
		bkp.compartments[index][item]++
	}

	return
}

func (bkp *backpack) getDuplicatePriority() int {
	if len(bkp.duplicates) != 1 {
		log.Fatalf("Did not find correct number of duplicates in %s, found %d", bkp.original, len(bkp.duplicates))
	}

	for key := range bkp.duplicates {
		fmt.Println(string(key), getItemPriority(key))
		return getItemPriority(key)
	}

	return -1000
}

func getItemPriority(item byte) int {
	if item >= 'a' && item <= 'z' {
		return int(item-'a') + 1
	}

	if item >= 'A' && item <= 'Z' {
		return int(item-'A') + 27
	}

	log.Fatalf("Trying to determine invalid item priority: %s", string(item))
	return -1
}
