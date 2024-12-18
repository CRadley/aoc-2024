package day07

import (
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Equation struct {
	Target int
	Values []int
}

func ParseIntegers(raw string) Equation {
	pattern, _ := regexp.Compile(`(\d+)`)
	matches := pattern.FindAllString(raw, -1)
	values := []int{}
	for _, m := range matches {
		v, _ := strconv.Atoi(m)
		values = append(values, v)
	}
	return Equation{Target: values[0], Values: values[1:]}
}

func Zfill(s string, t int) string {
	d := t - len(s)
	if d == 0 {
		return s
	}
	n := ""
	for range d {
		n += "0"
	}
	n += s
	return n
}

func GeneratePermutations(numValues int, base int) []string {
	permutations := []string{}
	actual := numValues - 1
	max := math.Pow(float64(base), float64(actual))
	for i := range int(max) {
		lenMax := len(strconv.FormatInt(int64(max), base))
		x := Zfill(strconv.FormatInt(int64(i), base), lenMax-1)
		permutations = append(permutations, x)
	}
	return permutations
}

func Parse(data []string) []Equation {
	lines := []Equation{}
	for _, line := range data {
		lines = append(lines, ParseIntegers(line))
	}
	return lines
}

func Execute(filepath string) (int, int, error) {
	raw, _ := os.ReadFile(filepath)
	data := strings.Split(string(raw), "\n")
	p1 := 0
	p2 := 0
	equations := Parse(data)
	eq2 := []Equation{}
	for _, e := range equations {
		ps := GeneratePermutations(len(e.Values), 2)
		found := false
		for _, p := range ps {
			t := 0
			t += e.Values[0]
			for i, s := range p {
				if string(s) == "0" {
					t += e.Values[i+1]
				} else {
					t *= e.Values[i+1]
				}
			}
			if t == e.Target {
				p1 += t
				p2 += t
				found = true
				break
			}
		}
		if !found {
			eq2 = append(eq2, e)
		}
	}
	for _, e := range eq2 {
		ps := GeneratePermutations(len(e.Values), 3)
		for _, p := range ps {
			t := 0
			t += e.Values[0]
			for i, s := range p {
				if string(s) == "0" {
					t += e.Values[i+1]
				} else if string(s) == "1" {
					x, _ := strconv.Atoi(strconv.Itoa(t) + strconv.Itoa(e.Values[i+1]))
					t = x
				} else {
					t *= e.Values[i+1]
				}
			}
			if t == e.Target {
				p2 += t
				break
			}
		}
	}
	return p1, p2, nil
}
