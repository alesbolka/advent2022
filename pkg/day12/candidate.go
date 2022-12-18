package day12

import "fmt"

type candidate struct {
	*coordinate
	distance int
}

func (cd candidate) String() string {
	return fmt.Sprintf("X: %d Y: %d H: %d Dist: %d", cd.x, cd.y, cd.height, cd.distance)
}
