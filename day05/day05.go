package day05

import (
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ParseManuals(rawManuals string) [][]int {
	manuals := [][]int{}
	for _, line := range strings.Split(rawManuals, "\n") {
		values := []int{}
		for _, v := range strings.Split(line, ",") {
			x, _ := strconv.Atoi(v)
			values = append(values, x)
		}
		manuals = append(manuals, values[:])

	}
	return manuals
}

func ParseRules(rawRules string) map[int][]int {
	var rules2 = map[int][]int{}
	for _, line := range strings.Split(rawRules, "\n") {
		s := strings.Split(line, "|")
		l, _ := strconv.Atoi(s[0])
		r, _ := strconv.Atoi(s[1])
		_, ok := rules2[l]
		if !ok {
			rules2[l] = []int{}
		}
		rules2[l] = append(rules2[l], r)
	}
	return rules2
}

func isValidManual(rules map[int][]int, m []int) bool {
	for i, value := range m {
		for _, j := range m[i+1:] {
			if !slices.Contains(rules[value], j) {
				return false
			}
		}
	}
	return true
}

func rearrange(rules map[int][]int, m *[]int) {
	for i := range *m {
		for j := range *m {
			if i == j {
				continue
			} else if !slices.Contains(rules[(*m)[j]], (*m)[i]) {
				(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
			}
		}
	}
}

func Execute(filepath string) (int, int, error) {
	data, _ := os.ReadFile(filepath)
	raw := string(data)
	splitInput := strings.Split(raw, "\n\n")
	rules := ParseRules(splitInput[0])
	manuals := ParseManuals(splitInput[1])
	p1 := 0
	p2 := 0
	for _, m := range manuals {
		midpoint := int(math.Floor(float64(len(m)) / float64(2)))
		if isValidManual(rules, m) {
			p1 += m[midpoint]
		} else {
			rearrange(rules, &m)
			p2 += m[midpoint]
		}
	}
	return p1, p2, nil
}
