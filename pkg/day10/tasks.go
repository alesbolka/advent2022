package day10

import "log"

func Task1(input string) int {
	cpu := NewCPU()
	res, err := cpu.RunProgram(input)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
