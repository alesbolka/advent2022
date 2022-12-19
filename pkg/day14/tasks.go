package day14

func Task1(input string) int {
	pit := NewSandpit(500, 0, input)

	return pit.TimeToVoid()
}

func Task2(input string) int {
	return 0
}
