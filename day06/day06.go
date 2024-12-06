package day06

import (
	"aoc2024/utils"
	"os"
	"slices"
)

func DetermineStartingPosition(lines [][]string) (int, int) {
	for i, r := range lines {
		for j, c := range r {
			if c == "^" {
				return i, j
			}
		}
	}
	return -1, -1
}

func Rotate(di, dj int) (int, int) {
	if di == -1 && dj == 0 {
		return 0, 1
	} else if di == 0 && dj == 1 {
		return 1, 0
	} else if di == 1 && dj == 0 {
		return 0, -1
	}
	return -1, 0
}

type Position struct {
	I int
	J int
}

type Position2 struct {
	I  int
	J  int
	DI int
	DJ int
}

func DetermineDir(i, j, pi, pj int) string {
	if i-pi == -1 {
		return "U"
	} else if i-pi == 1 {
		return "D"
	} else if j-pj == -1 {
		return "L"
	}
	return "R"
}

func Part2(lines [][]string, positions []Position) int {
	count := 0
	for _, p := range positions {
		i := p.I
		j := p.J
		if lines[i][j] == "^" {
			continue
		}
		lines[i][j] = "#"
		positions := []Position2{}
		ci, cj := DetermineStartingPosition(lines)
		di := -1
		dj := 0
		for {
			c := Position2{ci, cj, di, dj}
			if !slices.Contains(positions, c) {
				positions = append(positions, c)
			} else {
				count += 1
				break
			}
			if !utils.IsWithinBounds(ci+di, cj+dj, lines) {
				break
			}
			if lines[ci+di][cj+dj] == "#" {
				di, dj = Rotate(di, dj)
			} else {
				ci += di
				cj += dj
			}
		}
		lines[i][j] = "."
	}
	return count
}

func Execute(filepath string) (int, int, error) {
	raw, _ := os.ReadFile(filepath)
	lines := [][]string{}
	current := []string{}
	for _, c := range raw {
		if c == byte(10) {
			lines = append(lines, current[:])
			current = []string{}
		} else {
			current = append(current, string(c))
		}
	}
	lines = append(lines, current[:])
	ci, cj := DetermineStartingPosition(lines)
	di := -1
	dj := 0
	positions := []Position{}
	for {
		c := Position{ci, cj}
		if !slices.Contains(positions, c) {
			positions = append(positions, c)
		}
		if !utils.IsWithinBounds(ci+di, cj+dj, lines) {
			break
		}
		if lines[ci+di][cj+dj] == "#" {
			di, dj = Rotate(di, dj)
		} else {
			ci += di
			cj += dj
		}
	}
	return len(positions), Part2(lines, positions), nil
}
