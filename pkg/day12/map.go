package day12

import (
	"strings"
)

type Map struct {
	candidates []candidate
	end        *coordinate
	grid       [][]*coordinate
	start      *coordinate
}

var dirPairs [][2]int = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func NewMap(input string) *Map {
	res := &Map{}

	lines := strings.Split(input, "\n")
	res.grid = make([][]*coordinate, len(lines))

	for yy, line := range lines {
		res.grid[yy] = make([]*coordinate, len(line))
		for xx := 0; xx < len(line); xx++ {
			res.grid[yy][xx] = &coordinate{
				x:      xx,
				y:      yy,
				height: line[xx],
			}

			if line[xx] == 'S' {
				res.start = res.grid[yy][xx]
				res.start.height = 'a'
			} else if line[xx] == 'E' {
				res.end = res.grid[yy][xx]
				res.end.height = 'z'
			}
		}
	}

	return res
}

func (mp *Map) FindShortestPath() int {
	mp.candidates = []candidate{
		{
			coordinate: mp.start,
			distance:   0,
		},
	}

	mp.start.used = true

	for len(mp.candidates) != 0 {
		winner := mp.iterate()
		if winner != nil {
			return winner.distance
		}
	}
	return 0
}

func (mp *Map) iterate() *candidate {
	var cand candidate
	cand, mp.candidates = mp.candidates[0], mp.candidates[1:]

	for _, pair := range dirPairs {
		yy := cand.y + pair[0]
		if yy < 0 || yy >= len(mp.grid) {
			continue
		}

		xx := cand.x + pair[1]
		if xx < 0 || xx >= len(mp.grid[yy]) {
			continue
		}

		neighbour := mp.grid[yy][xx]

		if neighbour.used || neighbour.height > cand.height && (neighbour.height-cand.height) > 1 {
			continue
		}

		if neighbour == mp.end {
			return &candidate{
				coordinate: neighbour,
				distance:   cand.distance + 1,
			}
		}

		mp.insertCandidate(neighbour, cand.distance+1)
	}

	return nil
}

func (mp *Map) insertCandidate(node *coordinate, newDist int) {
	node.used = true
	lastIndex := len(mp.candidates) - 1
	newCand := candidate{
		coordinate: node,
		distance:   newDist,
	}

	if lastIndex < 0 || mp.candidates[lastIndex].distance <= newCand.distance {
		mp.candidates = append(mp.candidates, newCand)
		return
	}

	for ii := len(mp.candidates) - 1; ii >= 0; ii-- {
		if mp.candidates[ii].distance < newCand.distance {
			mp.candidates = append(mp.candidates[:ii+1], mp.candidates[ii:]...)
			mp.candidates[ii] = newCand
			return
		}
	}
}
