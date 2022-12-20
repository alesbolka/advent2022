package day15

import "fmt"

type coordinates struct {
	x int
	y int
}

func (cx coordinates) String() string {
	return fmt.Sprintf("(%d,%d)", cx.x, cx.y)
}

func (cx coordinates) InZone(min, max int) bool {
	return cx.x >= min && cx.x <= max &&
		cx.y >= min && cx.y <= max
}
