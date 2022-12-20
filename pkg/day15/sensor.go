package day15

type sensor struct {
	self        coordinates
	closest     coordinates
	emptyRadius int
}

func (sn sensor) InRange(other coordinates) bool {
	return manhattan(sn.self, other) <= sn.emptyRadius
}
