package day07

import (
	"strings"
)

func Task1(input string) (res uint64) {
	sys := newElfFs()

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		if line[0] == '$' {
			sys.executeCommand(line)
			continue
		}

		if sys.lsMode {
			sys.handleLSOutput(line)
		}
	}

	resMap := sys.getCombinedSmallFolderSize(100000)

	for _, size := range resMap {
		res += size
	}

	return
}

func Task2(input string) (res uint64) {
	sys := newElfFs()

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		if line[0] == '$' {
			sys.executeCommand(line)
			continue
		}

		if sys.lsMode {
			sys.handleLSOutput(line)
		}
	}

	// Just grab all
	minimumActual := uint64(70000000)
	resMap := sys.getCombinedSmallFolderSize(minimumActual)
	roomToMake := uint64(30000000) - (minimumActual - resMap["/"])

	for _, size := range resMap {
		if size > roomToMake && size < minimumActual {
			minimumActual = size
		}
	}
	return minimumActual
}
