package day09

import (
	"fmt"
	"strconv"
	"strings"
)

type KnotSystem struct {
	knots []struct {
		x int
		y int
	}
	report struct {
		x0 int
		x1 int
		y0 int
		y1 int
	}
	tailPositions map[string]bool
}

func newKnotSystem(knots int) (sys *KnotSystem) {
	sys = &KnotSystem{
		knots: make(
			[]struct {
				x int
				y int
			},
			knots,
		),
		report: struct {
			x0 int
			x1 int
			y0 int
			y1 int
		}{
			x0: 0,
			x1: 6,
			y0: 0,
			y1: 5,
		},
		tailPositions: map[string]bool{"0,0": true},
	}
	return
}

func (sys *KnotSystem) ExecuteMove(command string) error {
	args := strings.Split(command, " ")
	if len(args) != 2 {
		return fmt.Errorf("invalid command '%s'", command)
	}

	steps, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	knotLength := len(sys.knots)

	for step := 0; step < steps; step++ {
		switch args[0] {
		case "L":
			sys.knots[0].x--
		case "R":
			sys.knots[0].x++
		case "D":
			sys.knots[0].y--
		case "U":
			sys.knots[0].y++
		default:
			return fmt.Errorf("invalid direction in command: '%s'", command)
		}

		for ii := 1; ii < knotLength; ii++ {
			sys.knotMove(ii)
		}

		sys.tailPositions[fmt.Sprintf("%d,%d", sys.knots[knotLength-1].y, sys.knots[knotLength-1].x)] = true
	}

	return nil
}

func (sys *KnotSystem) knotMove(index int) {
	dX := sys.knots[index-1].x - sys.knots[index].x
	dY := sys.knots[index-1].y - sys.knots[index].y

	if dX > 1 {
		sys.knots[index].x++
		if dY == 1 || dY == -1 {
			sys.knots[index].y = sys.knots[index-1].y
		}
	} else if dX < -1 {
		sys.knots[index].x--
		if dY == 1 || dY == -1 {
			sys.knots[index].y = sys.knots[index-1].y
		}
	}

	if dY > 1 {
		sys.knots[index].y++
		if dX == 1 || dX == -1 {
			sys.knots[index].x = sys.knots[index-1].x
		}
	} else if dY < -1 {
		sys.knots[index].y--
		if dX == 1 || dX == -1 {
			sys.knots[index].x = sys.knots[index-1].x
		}
	}
}

func (sys *KnotSystem) String() (res string) {
	dY := sys.report.y1 - sys.report.y0 + 1
	dX := sys.report.x1 - sys.report.x0 + 1
	grid := make([][]byte, dY)

	for yy := dY - 1; yy >= 0; yy-- {
		grid[yy] = make([]byte, dX)
		for xx := 0; xx < dX; xx++ {
			grid[yy][xx] = '.'
		}
	}

	for ii, knot := range sys.knots {
		yy := knot.y - sys.report.y0
		xx := knot.x - sys.report.x0

		if grid[yy][xx] == '.' {
			grid[yy][xx] = '0' + byte(ii)

			if ii == 0 {
				grid[yy][xx] = 'H'
			}
		}
	}

	grid[-sys.report.y0][-sys.report.x0] = 's'

	for ii := dY - 1; ii >= 0; ii-- {
		fmt.Println(string(grid[ii][:]))
	}
	return
}
