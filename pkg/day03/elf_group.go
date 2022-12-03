package day03

import (
	"log"
)

type elfGroup struct {
	packs [3]*backpack
}

func (grp *elfGroup) findBadgePriority() int {
CHARACTER_LOOP:
	for item := range grp.packs[0].global {
		for otherBackpack := 1; otherBackpack < 3; otherBackpack++ {
			if _, ok := grp.packs[otherBackpack].global[item]; !ok {
				continue CHARACTER_LOOP
			}
		}

		return getItemPriority(item)
	}

	log.Fatalf("Failed to find a match, %s %s %s", grp.packs[0].original, grp.packs[1].original, grp.packs[2].original)
	return 0
}
