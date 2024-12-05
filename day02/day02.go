package day02

import (
	"aoc2024/utils"
	"os"
	"strconv"
	"strings"
)

func Abs(x, y int) int {
	value := x - y
	if value < 0 {
		return -value
	}
	return value
}
func IsAscending(values []int) bool {
	for i, v := range values[1:] {
		if values[i] >= v {
			return false
		}
	}
	return true
}

func IsDescending(values []int) bool {
	for i, v := range values[1:] {
		if values[i] <= v {
			return false
		}
	}
	return true
}

func AllWithinRange(values []int) bool {
	for i, v := range values[1:] {
		x := utils.Abs(values[i], v)
		if (1 > x) || (3 < x) {
			return false
		}
	}
	return true
}

func CanRemoveOne(value []int) bool {
	for i := range value {
		n := []int{}
		n = append(n, value[:i]...)
		n = append(n, value[i+1:]...)
		if (IsAscending(n) || IsDescending(n)) && AllWithinRange(n) {
			return true
		}
	}
	return false
}

func Execute(filepath string) (int, int, error) {
	data, _ := os.ReadFile(filepath)
	raw := strings.Split(string(data), "\n")
	lines := [][]int{}
	for _, line := range raw {
		s := []int{}
		r := strings.Split(string(line), " ")
		for _, _r := range r {
			x, _ := strconv.Atoi(_r)
			s = append(s, x)
		}
		lines = append(lines, s)
	}
	t1 := 0
	t2 := 0
	for _, l := range lines {
		if !IsDescending(l) && !IsAscending(l) {
			if CanRemoveOne(l) {
				t2 += 1
			}
		} else if AllWithinRange(l) {
			t1 += 1
			t2 += 1
		} else if CanRemoveOne(l) {
			t2 += 1
		}
	}
	return t1, t2, nil
}
