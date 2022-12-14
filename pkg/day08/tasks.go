package day08

import (
	"strings"
)

func Task1(input string) (res int) {
	lines := strings.Split(input, "\n")
	dimX := len(lines[0])
	dimY := len(lines)

	visibles := make([][]bool, dimY)
	startValue := byte('0') - 1

	tops := make([]byte, dimX)
	bots := make([]byte, dimX)
	for ii := 0; ii < dimX; ii++ {
		tops[ii] = startValue
		bots[ii] = startValue
	}
	for ii := 0; ii < dimY; ii++ {
		visibles[ii] = make([]bool, dimX)
	}

	for ii := 0; ii < dimY; ii++ {
		line := lines[ii]

		left := startValue
		right := startValue

		for jj := 0; jj < dimX; jj++ {
			if line[jj] > left {
				left = line[jj]
				if !visibles[ii][jj] {
					visibles[ii][jj] = true
					res++
				}
				// fmt.Println("left", line[jj]-'0')
			}

			if line[dimX-1-jj] > right {
				right = line[dimX-1-jj]
				if !visibles[ii][dimX-1-jj] {
					visibles[ii][dimX-1-jj] = true
					res++
				}
				// fmt.Println("right", line[lineLength-1-jj]-'0')
			}

			if line[jj] > tops[jj] {
				tops[jj] = line[jj]
				if !visibles[ii][jj] {
					visibles[ii][jj] = true
					res++
				}
			}

			botIndex := dimY - ii - 1
			if lines[botIndex][jj] > bots[jj] {
				bots[jj] = lines[botIndex][jj]
				if !visibles[botIndex][jj] {
					visibles[botIndex][jj] = true
					res++
				}
			}
		}
	}

	return
}

func Task2(input string) (res int) {

	lines := strings.Split(input, "\n")
	dimY := len(lines)
	dimX := len(lines[0])

	for houseY := 1; houseY < dimY-1; houseY++ {
		for houseX := 1; houseX < dimX-1; houseX++ {
			current := 1
			tempScore := 0

			// above
			for yy := houseY - 1; yy >= 0; yy-- {
				tempScore++

				if lines[yy][houseX] >= lines[houseY][houseX] {
					break
				}
			}

			current *= tempScore
			tempScore = 0

			// below
			for yy := houseY + 1; yy < dimY; yy++ {
				tempScore++

				if lines[yy][houseX] >= lines[houseY][houseX] {
					break
				}
			}

			current *= tempScore
			tempScore = 0

			// left
			for xx := houseX - 1; xx >= 0; xx-- {
				tempScore++

				if lines[houseY][xx] >= lines[houseY][houseX] {
					break
				}
			}

			current *= tempScore
			tempScore = 0

			// right
			for xx := houseX + 1; xx < dimX; xx++ {
				tempScore++

				if lines[houseY][xx] >= lines[houseY][houseX] {
					break
				}
			}

			current *= tempScore
			if current > res {
				res = current
			}
		}
	}
	return
}
