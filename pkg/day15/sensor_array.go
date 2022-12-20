package day15

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type SensorArray struct {
	sensors []sensor
	beacons map[coordinates]bool
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
		beacons: map[coordinates]bool{},
	}

	lineRegex := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)
	for _, line := range strings.Split(input, "\n") {
		match := lineRegex.FindStringSubmatch(line)
		if len(match) != 5 {
			return nil, fmt.Errorf("could not parse line '%s'", line)
		}

		sens := sensor{
			self:    coordinates{},
			closest: coordinates{},
		}
		sens.self.x, _ = strconv.Atoi(match[1])
		sens.self.y, _ = strconv.Atoi(match[2])
		sens.closest.x, _ = strconv.Atoi(match[3])
		sens.closest.y, _ = strconv.Atoi(match[4])
		res.beacons[sens.closest] = true
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

func (arr SensorArray) FindEmptySpot(zoneMin, zoneMax int) (res map[coordinates]bool) {
	res = map[coordinates]bool{}
	alreadyChecked := map[coordinates]bool{}

	for sensorIndex, sensor := range arr.sensors {
		// the beacon will be 1 off the radius, always
		offset := sensor.emptyRadius + 1

		candidates := []coordinates{}
		for ii := 0; ii <= 2*offset; ii++ {
			top := coordinates{sensor.self.x - offset + ii, sensor.self.y + ii}
			bottom := coordinates{sensor.self.x - offset + ii, sensor.self.y - ii}
			if top.x > sensor.self.x {
				top.y = sensor.self.y + 2*offset - ii
				bottom.y = sensor.self.y - 2*offset + ii
			}

			if top.InZone(zoneMin, zoneMax) && !alreadyChecked[top] {
				candidates = append(candidates, top)
			}

			if alreadyChecked[bottom] || top.y == sensor.self.y || !bottom.InZone(zoneMin, zoneMax) {
				continue
			}
			candidates = append(candidates, bottom)
		}

	CANDIDATE_LOOP:
		for _, candidate := range candidates {
			if arr.beacons[candidate] {
				continue
			}
			for checkIndex, otherSensor := range arr.sensors {
				if checkIndex == sensorIndex {
					continue
				}
				if otherSensor.InRange(candidate) {
					continue CANDIDATE_LOOP
				}
			}
			res[candidate] = true
		}
	}
	return
}
