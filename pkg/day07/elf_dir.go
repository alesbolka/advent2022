package day07

type elfDir struct {
	id      string
	parent  *elfDir
	dirSize uint64
	files   map[string]uint64
	dirs    map[string]*elfDir
}

func (dir *elfDir) getSize(threshold uint64, resMap map[string]uint64) uint64 {
	if dir.dirSize > 0 {
		return dir.dirSize
	}

	for _, subDir := range dir.dirs {
		dir.dirSize += subDir.getSize(threshold, resMap)
	}

	for _, size := range dir.files {
		dir.dirSize += size
	}

	if dir.dirSize < threshold {
		resMap[dir.id] = dir.dirSize
	}

	return dir.dirSize
}
