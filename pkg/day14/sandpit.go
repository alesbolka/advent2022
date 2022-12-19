package day14

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type coordinate struct {
	X int
	Y int
}

type Sandpit struct {
	source coordinate
	room   map[int]map[int]byte
	rangeX [2]int
	rangeY [2]int
}

func NewSandpit(sourceX, sourceY int, input string) *Sandpit {
	pit := &Sandpit{
		source: coordinate{X: sourceX, Y: sourceY},
		room: map[int]map[int]byte{
			sourceY: {
				sourceX: 'S',
			},
		},
		rangeX: [2]int{sourceX, sourceX},
		rangeY: [2]int{sourceY, sourceY},
	}

	for _, line := range strings.Split(input, "\n") {
		firstRun := true
		var currentX, currentY int

		for _, corner := range strings.Split(line, " -> ") {
			targetCoords := strings.Split(corner, ",")

			targetY, err := strconv.Atoi(targetCoords[1])
			if err != nil {
				log.Fatalf("Failed parsing number from '%s' at pair '%s'", line, corner)
			}

			targetX, err := strconv.Atoi(targetCoords[0])
			if err != nil {
				log.Fatalf("Failed parsing number from '%s' at pair '%s'", line, corner)
			}

			if firstRun {
				firstRun = false
				currentX = targetX
				currentY = targetY
				continue
			}

			stepY := 1
			stepX := 1
			if currentY > targetY {
				stepY = -1
			} else if currentY == targetY {
				stepY = 0
			}

			if currentX > targetX {
				stepX = -1
			} else if currentX == targetX {
				stepX = 0
			}

			for currentX != targetX || currentY != targetY {
				if _, exists := pit.room[currentY]; !exists {
					pit.room[currentY] = map[int]byte{}
				}
				pit.room[currentY][currentX] = '#'

				if currentX < pit.rangeX[0] {
					pit.rangeX[0] = currentX
				}
				if currentX > pit.rangeX[1] {
					pit.rangeX[1] = currentX
				}

				if currentY < pit.rangeY[0] {
					pit.rangeY[0] = currentY
				}
				if currentY > pit.rangeY[1] {
					pit.rangeY[1] = currentY
				}
				currentX += stepX
				currentY += stepY
			}

			if _, exists := pit.room[currentY]; !exists {
				pit.room[currentY] = map[int]byte{}
			}
			pit.room[currentY][currentX] = '#'
		}
	}

	return pit
}

func (pit *Sandpit) emptySand() {
	pit.room[pit.source.Y][pit.source.X] = 'S'
	for yy := range pit.room {
		for xx, val := range pit.room[yy] {
			if val == 2 {
				pit.room[yy][xx] = 0
			}
		}
	}
}

func (pit *Sandpit) TimeToVoid() int {
	pit.emptySand()
	for ii := 0; ii < 1000000; ii++ {
		stopDepth := pit.dropGrain()
		// fmt.Println("Grain", ii, stopDepth)
		// fmt.Println(pit)
		if stopDepth > pit.rangeY[1] {
			return ii
		}
	}
	return -1
}

func (pit *Sandpit) FillUp() (res int) {
	pit.emptySand()
	for pit.room[pit.source.Y][pit.source.X] == 'S' {
		pit.dropGrain()
		res++
	}
	return
}

func (pit *Sandpit) dropGrain() int {
	xx, yy := pit.source.X, pit.source.Y

	for {
		if yy == pit.rangeY[1]+1 {
			if _, exists := pit.room[yy]; !exists {
				pit.room[yy] = map[int]byte{}
			}
			pit.room[yy][xx] = 'o'
			return yy
		}

		if pit.room[yy+1][xx] == 0 {
			yy++
			continue
		}

		belowLeft := pit.room[yy+1][xx-1]
		if belowLeft == 0 {
			yy++
			xx--
			continue
		}

		belowRight := pit.room[yy+1][xx+1]
		if belowRight == 0 {
			yy++
			xx++
			continue
		}

		if _, exists := pit.room[yy]; !exists {
			pit.room[yy] = map[int]byte{}
		}
		pit.room[yy][xx] = 'o'
		return yy
	}
}

func (pit *Sandpit) String() (res string) {
	ground := pit.rangeY[1] + 2
	for yy := pit.rangeY[0]; yy <= ground; yy++ {
		for xx := pit.rangeX[0] - 4; xx <= pit.rangeX[1]+3; xx++ {
			if yy == ground {
				res += "#"
				continue
			}
			val := pit.room[yy][xx]
			switch val {
			case 0:
				res += "."
			case '#':
				fallthrough
			case 'o':
				fallthrough
			case 'S':
				res += string(val)
			default:
				fmt.Printf("Invalid map value %d at %d-%d\n", val, yy, xx)
				res += "!"
			}
		}
		res += "\n"
	}
	return
}
