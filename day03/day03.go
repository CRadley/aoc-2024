package day03

import (
	"os"
	"regexp"
	"strconv"
)

func Execute(filepath string) (int, int, error) {
	data, _ := os.ReadFile(filepath)
	raw := string(data)
	pattern, _ := regexp.Compile(`mul\([\d]{1,3},[\d]{1,3}\)|do\(\)|don't\(\)`)
	pattern2, _ := regexp.Compile(`\d+`)
	matches := pattern.FindAllString(raw, -1)

	shouldAdd := true
	t := 0
	t2 := 0
	for _, m := range matches {
		switch m {
		case "do()":
			shouldAdd = true
		case "don't()":
			shouldAdd = false
		default:
			values := pattern2.FindAllString(m, -1)
			v1, _ := strconv.Atoi(values[0])
			v2, _ := strconv.Atoi(values[1])
			t += (v1 * v2)
			if shouldAdd {
				t2 += (v1 * v2)
			}
		}

	}
	return t, t2, nil
}
