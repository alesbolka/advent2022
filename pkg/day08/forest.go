package day08

import "strings"

type forest struct {
}

func parseMap(aerialMap string) *forest {
	lines := strings.Split(aerialMap, "\n")
	grid := make([][]int, len(lines))

	for ii := 0; ii < len(lines); ii++ {
		lineLength := len(lines[ii])
		grid[ii] = make([]int, lineLength)

		for jj := 0; jj < lineLength; jj++ {
			grid[ii][jj] = int(lines[ii][jj] - '0')
		}
	}

	return nil
}
