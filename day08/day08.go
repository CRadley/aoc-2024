package day08

import (
	"aoc2024/utils"
	"os"
	"slices"
	"strings"
)

type Point struct {
	I int
	J int
}

type Pair struct {
	A Point
	B Point
}

func (p Pair) Diff() (int, int) {
	return p.B.I - p.A.I, p.B.J - p.A.J
}

func DeterminePairs(points []Point) []Pair {
	pairs := []Pair{}
	for i, point := range points {
		for _, point2 := range points[i+1:] {
			pairs = append(pairs, Pair{point, point2})
		}
	}
	return pairs
}

func ParseRoof(raw []string) map[string][]Point {
	r := map[string][]Point{}
	for i, row := range raw {
		for j, value := range row {
			s := string(value)
			if s == "." {
				continue
			}
			_, found := r[s]
			if !found {
				r[s] = []Point{}
			}
			r[s] = append(r[s], Point{i, j})
		}
	}
	return r
}

func Execute(filepath string) (int, int, error) {
	raw, _ := os.ReadFile(filepath)
	s := strings.Split(string(raw), "\n")
	roof := ParseRoof(s)
	x := []Point{}
	y := []Point{}
	for _, v := range roof {
		for _, p := range DeterminePairs(v) {
			di, dj := p.Diff()
			i2 := p.A.I
			j2 := p.A.J
			i3 := p.B.I
			j3 := p.B.J
			y = append(y, p.A)
			y = append(y, p.B)
			x = append(x, Point{p.A.I - di, p.A.J - dj})
			x = append(x, Point{p.B.I + di, p.B.J + dj})
			for {
				i2 -= di
				j2 -= dj
				if !utils.IsWithinBounds2(i2, j2, s) {
					break
				}
				y = append(y, Point{i2, j2})
			}
			for {
				i3 += di
				j3 += dj
				if !utils.IsWithinBounds2(i3, j3, s) {
					break
				}
				y = append(y, Point{i3, j3})
			}
		}
	}
	actual := []Point{}
	for _, p := range x {
		if !slices.Contains(actual, p) && utils.IsWithinBounds2(p.I, p.J, s) {
			actual = append(actual, p)
		}
	}
	ay := []Point{}
	for _, p := range y {
		if !slices.Contains(ay, p) {
			ay = append(ay, p)
		}
	}
	p1 := len(actual)
	p2 := len(ay)
	return p1, p2, nil
}
