package day09

import (
	"fmt"
	"strconv"
	"strings"
)

type RopeSystem struct {
	head struct {
		xx int
		yy int
	}
	tail struct {
		xx int
		yy int
	}
	printRange struct {
		x0 int
		x1 int
		y0 int
		y1 int
	}
	visited        map[int]map[int]bool
	VisitedCounter uint64
}

func NewSystem(startX, startY int) *RopeSystem {
	return &RopeSystem{
		head: struct {
			xx int
			yy int
		}{
			xx: startX,
			yy: startY,
		},
		tail: struct {
			xx int
			yy int
		}{
			xx: startX,
			yy: startY,
		},
		printRange: struct {
			x0 int
			x1 int
			y0 int
			y1 int
		}{
			x1: 6,
			y1: 5,
		},
		visited: map[int]map[int]bool{
			startY: {
				startX: true,
			},
		},
		VisitedCounter: 1,
	}
}

func (sys *RopeSystem) ExecuteMove(command string) (err error) {
	args := strings.Split(command, " ")
	if len(args) != 2 {
		return fmt.Errorf("invalid command '%s'", command)
	}

	steps, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	var dX, dY int

	for ii := 0; ii < steps; ii++ {
		switch args[0] {
		case "L":
			sys.head.xx--
		case "R":
			sys.head.xx++
		case "D":
			sys.head.yy--
		case "U":
			sys.head.yy++
		default:
			return fmt.Errorf("invalid direction in command: '%s'", command)
		}

		dX = sys.head.xx - sys.tail.xx
		dY = sys.head.yy - sys.tail.yy
		if dX > 1 {
			sys.tail.xx++
			if dY == 1 || dY == -1 {
				sys.tail.yy = sys.head.yy
			}
		} else if dX < -1 {
			sys.tail.xx--
			if dY == 1 || dY == -1 {
				sys.tail.yy = sys.head.yy
			}
		}

		if dY > 1 {
			sys.tail.yy++
			if dX == 1 || dX == -1 {
				sys.tail.xx = sys.head.xx
			}
		} else if dY < -1 {
			sys.tail.yy--
			if dX == 1 || dX == -1 {
				sys.tail.xx = sys.head.xx
			}
		}

		if _, exists := sys.visited[sys.tail.yy]; !exists {
			sys.visited[sys.tail.yy] = map[int]bool{}
		}

		if _, exists := sys.visited[sys.tail.yy][sys.tail.xx]; !exists {
			sys.visited[sys.tail.yy][sys.tail.xx] = true
			sys.VisitedCounter++
		}
	}
	return
}

func (sys *RopeSystem) String() (res string) {
	for yy := sys.printRange.y1 - 1; yy >= sys.printRange.y0; yy-- {
		for xx := sys.printRange.x0; xx < sys.printRange.x1; xx++ {
			if xx == sys.head.xx && yy == sys.head.yy && xx == sys.tail.xx && yy == sys.tail.yy {
				res += "X"
			} else if xx == sys.head.xx && yy == sys.head.yy {
				res += "H"
			} else if xx == sys.tail.xx && yy == sys.tail.yy {
				res += "T"
			} else {
				res += "."
			}
		}
		res += "\n"
	}
	return
}
