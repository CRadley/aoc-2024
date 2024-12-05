package day04

import (
	"os"
	"reflect"
)

func IsWithinBounds(i, j int, lines [][]string) bool {
	return (i >= 0 && i < len(lines)) && (j >= 0 && j < len(lines[0]))
}

func Execute(filepath string) (int, int, error) {
	XMAS := []string{"X", "M", "A", "S"}
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
	count := 0
	count2 := 0
	for i, r := range lines {
		for j, value := range r {
			if value == "A" {
				if IsWithinBounds(i-1, j-1, lines) && IsWithinBounds(i+1, j-1, lines) && IsWithinBounds(i-1, j+1, lines) && IsWithinBounds(i+1, j+1, lines) {
					if lines[i-1][j-1] == "M" && lines[i-1][j+1] == "M" && lines[i+1][j-1] == "S" && lines[i+1][j+1] == "S" {
						count2 += 1
					}
					if lines[i-1][j-1] == "S" && lines[i-1][j+1] == "S" && lines[i+1][j-1] == "M" && lines[i+1][j+1] == "M" {
						count2 += 1
					}
					if lines[i-1][j-1] == "M" && lines[i-1][j+1] == "S" && lines[i+1][j-1] == "M" && lines[i+1][j+1] == "S" {
						count2 += 1
					}
					if lines[i-1][j-1] == "S" && lines[i-1][j+1] == "M" && lines[i+1][j-1] == "S" && lines[i+1][j+1] == "M" {
						count2 += 1
					}
				}
			} else if value == "X" {
				// RIGHT
				if IsWithinBounds(i, j+1, lines) && IsWithinBounds(i, j+2, lines) && IsWithinBounds(i, j+3, lines) {
					if reflect.DeepEqual([]string{value, r[j+1], r[j+2], r[j+3]}, XMAS) {
						count += 1
					}
				}
				// LEFT
				if IsWithinBounds(i, j-1, lines) && IsWithinBounds(i, j-2, lines) && IsWithinBounds(i, j-3, lines) {
					if reflect.DeepEqual([]string{value, r[j-1], r[j-2], r[j-3]}, XMAS) {
						count += 1
					}
				}
				// DOWN
				if IsWithinBounds(i+1, j, lines) && IsWithinBounds(i+2, j, lines) && IsWithinBounds(i+3, j, lines) {
					if reflect.DeepEqual([]string{value, lines[i+1][j], lines[i+2][j], lines[i+3][j]}, XMAS) {
						count += 1
					}
				}
				// UP
				if IsWithinBounds(i-1, j, lines) && IsWithinBounds(i-2, j, lines) && IsWithinBounds(i-3, j, lines) {
					if reflect.DeepEqual([]string{value, lines[i-1][j], lines[i-2][j], lines[i-3][j]}, XMAS) {
						count += 1
					}
				}
				// UP RIGHT
				if IsWithinBounds(i+1, j-1, lines) && IsWithinBounds(i+2, j-2, lines) && IsWithinBounds(i+3, j-3, lines) {
					if reflect.DeepEqual([]string{value, lines[i+1][j-1], lines[i+2][j-2], lines[i+3][j-3]}, XMAS) {
						count += 1
					}
				}
				// UP LEFT
				if IsWithinBounds(i-1, j-1, lines) && IsWithinBounds(i-2, j-2, lines) && IsWithinBounds(i-3, j-3, lines) {
					if reflect.DeepEqual([]string{value, lines[i-1][j-1], lines[i-2][j-2], lines[i-3][j-3]}, XMAS) {
						count += 1
					}
					// DOWN RIGHT
				}
				if IsWithinBounds(i+1, j+1, lines) && IsWithinBounds(i+2, j+2, lines) && IsWithinBounds(i+3, j+3, lines) {
					if reflect.DeepEqual([]string{value, lines[i+1][j+1], lines[i+2][j+2], lines[i+3][j+3]}, XMAS) {
						count += 1
					}
				}
				// DOWN LEFT
				if IsWithinBounds(i-1, j+1, lines) && IsWithinBounds(i-2, j+2, lines) && IsWithinBounds(i-3, j+3, lines) {
					if reflect.DeepEqual([]string{value, lines[i-1][j+1], lines[i-2][j+2], lines[i-3][j+3]}, XMAS) {
						count += 1
					}
				}
			}
		}
	}

	return count, count2, nil
}