package day07

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type elfSystem struct {
	lsMode bool
	root   *elfDir
	cwd    *elfDir
}

func newElfFs() *elfSystem {
	sys := &elfSystem{
		root: newDir("/", nil),
	}

	return sys
}

func newDir(id string, parent *elfDir) *elfDir {
	fullPath := id
	if parent != nil {
		fullPath = fmt.Sprintf("%s%s/", parent.id, id)
	}

	return &elfDir{
		id:     fullPath,
		parent: parent,
		files:  map[string]uint64{},
		dirs:   map[string]*elfDir{},
	}
}

func (sys *elfSystem) executeCommand(command string) {
	sys.lsMode = false
	args := strings.Split(command, " ")

	switch args[1] {
	case "cd":
		switch args[2] {
		case "..":
			if sys.cwd.parent == nil {
				// (should this even error?)
				log.Fatal("Cannot cd back from root")
			}
			sys.cwd = sys.cwd.parent
		case "/":
			sys.cwd = sys.root
		default:
			newCwd, exists := sys.cwd.dirs[args[2]]
			if !exists {
				log.Fatalf("cannot cd, no directory with such name, command: %s", command)
			}
			sys.cwd = newCwd
		}
	case "ls":
		sys.lsMode = true
	default:
		log.Fatalf("Invalid command: %s", command)
	}
}

func (sys *elfSystem) handleLSOutput(output string) {
	parts := strings.Split(output, " ")
	if len(parts) != 2 {
		log.Fatalf("could not parse ls output for line %s", output)
	}

	if parts[0] == "dir" {
		if _, exists := sys.cwd.dirs[parts[1]]; !exists {
			sys.cwd.dirs[parts[1]] = newDir(parts[1], sys.cwd)
		}
		return
	}

	val, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		log.Fatalf("Failed converting file size for %s", output)
	}

	sys.cwd.files[parts[1]] = val
}

func (sys *elfSystem) getCombinedSmallFolderSize(threshold uint64) (resMap map[string]uint64) {
	resMap = map[string]uint64{}
	sys.root.getSize(threshold, resMap)

	return
}
