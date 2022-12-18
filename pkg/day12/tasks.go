package day12

func Task1(input string) int {
	mp := NewMap(input)

	return mp.FindShortestPath(mp.Start)
}

func Task2(input string) int {
	mp := NewMap(input)
	starts := mp.GetAllOfElevation('a')
	minimum := 0

	for _, start := range starts {
		localRes := mp.FindShortestPath(start)
		if localRes == 0 {
			continue
		}
		if minimum == 0 || localRes < minimum {
			minimum = localRes
		}
	}

	return minimum
}
