package day07

import "strings"

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

	return sys.getCombinedSmallFolderSize(100000)
}
