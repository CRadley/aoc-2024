package day15

import (
	"os"
	"slices"
	"sort"
	"strings"
)

func GetStartingPosition(storageMap [][]string) (int, int) {
	for i, r := range storageMap {
		for j, v := range r {
			if v == "@" {
				return i, j
			}
		}
	}
	return -1, -1
}

type Direction struct {
	H int
	V int
}

var DIRECTIONS = map[string]Direction{
	"^": {0, -1},
	">": {1, 0},
	"v": {0, 1},
	"<": {-1, 0},
}

func CreateExpandedStorageMap(storageMap [][]string) [][]string {
	expandedStorageMap := make([][]string, len(storageMap))
	for i, row := range storageMap {
		expandedStorageMap[i] = make([]string, len(storageMap[0])*2)
		for j, v := range row {
			if v == "#" {
				expandedStorageMap[i][2*j] = "#"
				expandedStorageMap[i][2*j+1] = "#"
			} else if v == "." {
				expandedStorageMap[i][2*j] = "."
				expandedStorageMap[i][2*j+1] = "."
			} else if v == "@" {
				expandedStorageMap[i][2*j] = "@"
				expandedStorageMap[i][2*j+1] = "."
			} else if v == "O" {
				expandedStorageMap[i][2*j] = "["
				expandedStorageMap[i][2*j+1] = "]"
			}
		}
	}
	return expandedStorageMap
}

func CalculatePartOne(storageMap [][]string, moves string) int {
	i, j := GetStartingPosition(storageMap)
	for _, move := range moves {
		if string(move) == "\n" {
			continue
		}
		direction := DIRECTIONS[string(move)]
		ni := i + direction.V
		nj := j + direction.H
		if storageMap[ni][nj] == "#" {
			continue
		} else if storageMap[ni][nj] == "O" {
			nni, nnj := ni, nj
			for {
				nni += direction.V
				nnj += direction.H
				if storageMap[nni][nnj] == "#" || storageMap[nni][nnj] == "." {
					break
				}
			}
			if storageMap[nni][nnj] == "#" {
				continue
			}
			storageMap[ni][nj] = "@"
			storageMap[i][j] = "."
			i, j = ni, nj
			bi, bj := ni, nj
			for {
				bi += direction.V
				bj += direction.H
				storageMap[bi][bj] = "O"
				if bi == nni && bj == nnj {
					break
				}
			}
		} else {
			storageMap[ni][nj] = "@"
			storageMap[i][j] = "."
			i, j = ni, nj
		}
	}
	p1 := 0
	for i, r := range storageMap {
		for j, v := range r {
			if v == "O" {
				p1 += (100*i + j)
			}
		}
	}
	return p1
}

type Pos struct {
	I int
	J int
	C string
}

func CanPush(storageMap [][]string, dir Direction, i, j int, positions *[]Pos) bool {
	if !slices.Contains((*positions), Pos{i, j, storageMap[i][j]}) {
		(*positions) = append((*positions), Pos{i, j, storageMap[i][j]})
	}
	ni := i + dir.V
	nj := j + dir.H
	if storageMap[ni][nj] == "#" {
		return false
	} else if storageMap[ni][nj] == "." {
		return true
	}
	if dir == DIRECTIONS["<"] {
		if !CanPush(storageMap, dir, ni, nj, positions) {
			return false
		}
	} else if dir == DIRECTIONS[">"] {
		if !CanPush(storageMap, dir, ni, nj, positions) {
			return false
		}
	} else {
		if !CanPush(storageMap, dir, ni, nj, positions) {
			return false
		}
		if storageMap[ni][nj] == "]" {
			if !CanPush(storageMap, dir, ni, nj-1, positions) {
				return false
			}
		} else {
			if !CanPush(storageMap, dir, ni, nj+1, positions) {
				return false
			}
		}
	}
	return true
}

func CalculatePartTwo(storageMap [][]string, moves string) int {
	for _, move := range moves {
		if string(move) == "\n" {
			continue
		}
		direction := DIRECTIONS[string(move)]
		i, j := GetStartingPosition(storageMap)
		position := &[]Pos{}
		if CanPush(storageMap, direction, i, j, position) {
			if direction == DIRECTIONS[">"] || direction == DIRECTIONS["<"] {
				for k := len(*position) - 1; k >= 0; k-- {
					p := (*position)[k]
					storageMap[p.I][p.J], storageMap[p.I+direction.V][p.J+direction.H] = storageMap[p.I+direction.V][p.J+direction.H], storageMap[p.I][p.J]
				}
			} else {
				values := map[int][]int{}
				for _, k := range *position {
					_, ok := values[k.I]
					if !ok {
						values[k.I] = []int{}
					}
					values[k.I] = append(values[k.I], k.J)
				}
				keys := make([]int, len(values))
				i := 0
				for k := range values {
					keys[i] = k
					i++
				}
				slices.Sort(keys)
				if direction == DIRECTIONS["^"] {
					for _, s := range keys {
						slices.Sort(values[s])
						for _, v := range values[s] {
							storageMap[s][v], storageMap[s+direction.V][v+direction.H] = storageMap[s+direction.V][v+direction.H], storageMap[s][v]
						}
					}
				} else {
					sort.Sort(sort.Reverse(sort.IntSlice(keys)))
					for _, s := range keys {
						sort.Sort(sort.Reverse(sort.IntSlice(values[s])))
						for _, v := range values[s] {
							storageMap[s][v], storageMap[s+direction.V][v+direction.H] = storageMap[s+direction.V][v+direction.H], storageMap[s][v]
						}
					}
				}
			}
		}
	}
	total := 0
	for i, r := range storageMap {
		for j, v := range r {
			if v == "[" {
				total += (100*i + j)
			}
		}
	}
	return total
}

func Execute(filepath string) (int, int, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return -1, -1, err
	}
	data := string(bytes)
	split := strings.Split(data, "\n\n")
	rawMap, rawMoves := split[0], split[1]
	storageMap := strings.Split(rawMap, "\n")
	sm := make([][]string, len(storageMap))
	for i := range sm {
		sm[i] = make([]string, len(storageMap[0]))
		for j := range len(storageMap[0]) {
			sm[i][j] = string(storageMap[i][j])
		}
	}
	ex := CreateExpandedStorageMap(sm)
	return CalculatePartOne(sm, rawMoves), CalculatePartTwo(ex, rawMoves), nil
}
