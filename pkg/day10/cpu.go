package day10

import (
	"fmt"
	"strconv"
	"strings"
)

type CPU struct {
	currentCycle   int
	registry       int
	signalStrength int
	screen         [6][40]byte
}

func NewCPU() *CPU {
	return &CPU{
		currentCycle: 1,
		registry:     1,
	}
}

func (cpu *CPU) cycle() {
	displayRow := (cpu.currentCycle - 1) / 40
	displayColumn := (cpu.currentCycle - 1) % 40

	cpu.screen[displayRow][displayColumn] = '.'
	if displayColumn >= (cpu.registry-1) && displayColumn <= (cpu.registry+1) {
		cpu.screen[displayRow][displayColumn] = '#'
	}

	if cpu.currentCycle >= 20 && (cpu.currentCycle-20)%40 == 0 {
		cpu.signalStrength += cpu.currentCycle * cpu.registry
	}
	cpu.currentCycle++
}

func (cpu *CPU) noop() {
	cpu.cycle()
}

func (cpu *CPU) addx(val int) {
	cpu.cycle()
	cpu.cycle()
	cpu.registry += val
}

func (cpu *CPU) RunProgram(commands string) (res int, err error) {
	list := strings.Split(commands, "\n")

	for _, command := range list {
		commandArgs := strings.Split(command, " ")
		if len(commandArgs) != 2 && len(commandArgs) != 1 {
			return -1, fmt.Errorf("invalid command '%s'", command)
		}

		switch commandArgs[0] {
		case "noop":
			cpu.noop()
		case "addx":
			val, err := strconv.Atoi(commandArgs[1])
			if err != nil {
				return -1, fmt.Errorf("could not parse value from '%s'", command)
			}

			cpu.addx(val)
		default:
			return -1, fmt.Errorf("unreckognized command '%s'", command)
		}
	}

	return cpu.signalStrength, nil
}

func (cpu *CPU) ShowScreen() (res string) {
	for ii := 0; ii < 6; ii++ {
		res += string(cpu.screen[ii][:]) + "\n"
	}
	return
}
