package day12

type coordinate struct {
	x      int
	y      int
	height byte
	used   bool
}

func (aa coordinate) Equal(bb coordinate) bool {
	return aa.x == bb.x && aa.y == bb.y
}
