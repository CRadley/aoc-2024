package day01

import (
	"aoc2024/utils"
	"os"
	"slices"
	"strconv"
	"strings"
)

func DetermineCounts(values []int) map[int]int {
	counts := map[int]int{}
	for _, v := range values {
		_, exists := counts[v]
		if !exists {
			counts[v] = 0
		}
		counts[v] += 1
	}
	return counts
}

func parseInputFile(data []byte) ([]int, []int, error) {
	raw := strings.Split(string(data), "\n")
	left := []int{}
	right := []int{}
	for _, line := range raw {
		values := strings.Split(line, "   ")
		l, err := strconv.Atoi(values[0])
		if err != nil {
			return nil, nil, err
		}
		r, err := strconv.Atoi(values[1])
		if err != nil {
			return nil, nil, err
		}
		left = append(left, l)
		right = append(right, r)
	}
	return left, right, nil
}

func Execute(filepath string) (int, int, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return -1, -1, err
	}
	left, right, err := parseInputFile(data)
	if err != nil {
		return -1, -1, err
	}
	slices.Sort(left)
	slices.Sort(right)
	t1 := 0
	t2 := 0
	counts := DetermineCounts(right)
	for i, v := range left {
		t1 += utils.Abs(v, right[i])
		t2 += v * counts[v]
	}
	return t1, t2, nil
}
