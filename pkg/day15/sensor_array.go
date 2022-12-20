package day15

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type coordinates struct {
	x int
	y int
}

func (cx coordinates) String() string {
	return fmt.Sprintf("(%d,%d)", cx.x, cx.y)
}

type sensor struct {
	self        coordinates
	closest     coordinates
	emptyRadius int
}

type SensorArray struct {
	sensors []sensor
}

func manhattan(a, b coordinates) (res int) {
	if a.x < b.x {
		res = b.x - a.x
	} else {
		res = a.x - b.x
	}

	if a.y < b.y {
		res += b.y - a.y
	} else {
		res += a.y - b.y
	}

	return
}

func NewArray(input string) (*SensorArray, error) {
	res := &SensorArray{
		sensors: []sensor{},
	}

	lineRegex := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)
	for _, line := range strings.Split(input, "\n") {
		match := lineRegex.FindStringSubmatch(line)
		if len(match) != 5 {
			return nil, fmt.Errorf("could not parse line '%s'", line)
		}

		sens := sensor{}
		sens.self.x, _ = strconv.Atoi(match[1])
		sens.self.y, _ = strconv.Atoi(match[2])
		sens.closest.x, _ = strconv.Atoi(match[3])
		sens.closest.y, _ = strconv.Atoi(match[4])
		sens.emptyRadius = manhattan(sens.self, sens.closest)

		res.sensors = append(res.sensors, sens)
	}
	return res, nil
}

func (arr SensorArray) CountEmptySpots(targetRow int) (res int) {
	guaranteedEmpty := map[int]bool{}

	for _, sensor := range arr.sensors {
		rangeLimit := sensor.emptyRadius
		if sensor.closest.y == targetRow {
			guaranteedEmpty[sensor.closest.x] = false
		}

		dY := targetRow - sensor.self.y
		if dY < 0 {
			dY = sensor.self.y - targetRow
		}

		if dY > rangeLimit {
			// These are simply out of range for the target row
			continue
		}
		dX := rangeLimit - dY
		for xx := sensor.self.x - dX; xx <= sensor.self.x+dX; xx++ {
			if _, exists := guaranteedEmpty[xx]; !exists {
				res++
				guaranteedEmpty[xx] = true
			}
		}
	}

	return
}
