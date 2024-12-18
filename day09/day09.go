package day09

import (
	"os"
	"strconv"
)

type FileSystem []int

func (fs *FileSystem) Count(target int) int {
	c := 0
	for _, v := range *fs {
		if v == target {
			c += 1
		}
	}
	return c
}

func (fs *FileSystem) EmptyBlockSizeExists(size int) (int, bool) {
	inEmptyBlock := false
	ci := -1
	cs := 0
	for i, v := range *fs {
		if v == -1 {
			if inEmptyBlock {
				cs += 1
			} else {
				cs = 1
				ci = i
				inEmptyBlock = true
			}
			if cs == size {
				return ci, true
			}
		} else {
			cs = 0
			inEmptyBlock = false
		}
	}
	return -1, false
}

func (fs *FileSystem) GetTarget(target int) int {
	for i, v := range *fs {
		if v == target {
			return i
		}
	}
	return -1
}

func (fs FileSystem) String() string {
	s := ""
	for _, v := range fs {
		if v == -1 {
			s += "."
		} else {
			s += strconv.Itoa(v)
		}
	}
	return s
}

func (fs *FileSystem) Fragment() {
	l := 0
	r := len(*fs) - 1
	for {
		if l >= r {
			break
		}
		if (*fs)[l] == -1 {
			for {
				if l >= r {
					break
				}
				if (*fs)[r] != -1 {
					(*fs)[l], (*fs)[r] = (*fs)[r], (*fs)[l]
					break
				} else {
					r -= 1
				}
			}
		}
		l += 1
	}

}

func (fs *FileSystem) FileMove(currentId int) {
	currentId -= 1
	for {
		if currentId == -1 {
			break
		}
		blockSize := fs.Count(currentId)
		i, found := fs.EmptyBlockSizeExists(blockSize)
		s := fs.GetTarget(currentId)
		if found && i < s {
			for range blockSize {
				(*fs)[i], (*fs)[s] = (*fs)[s], (*fs)[i]
				i += 1
				s += 1
			}
		}
		currentId -= 1
	}
}

func (fs *FileSystem) Value() int {
	// Fragment
	t := 0
	for i, v := range *fs {
		if v == -1 {
			continue
		}
		t += i * v
	}
	return t
}

func Execute(filepath string) (int, int, error) {
	raw, _ := os.ReadFile(filepath)
	data := string(raw)
	currentId := 0
	fs := FileSystem{}
	fs2 := FileSystem{}
	for i, v := range data {
		x, _ := strconv.Atoi(string(v))
		for range x {
			if i%2 == 0 {
				fs = append(fs, currentId)
				fs2 = append(fs2, currentId)
			} else {
				fs = append(fs, -1)
				fs2 = append(fs2, -1)
			}
		}
		if i%2 == 0 {
			currentId += 1
		}
	}
	fs.Fragment()
	fs2.FileMove(currentId)
	return fs.Value(), fs2.Value(), nil
}
