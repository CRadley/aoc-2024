package day12

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	I int
	J int
}

const (
	U = iota
	R = iota
	D = iota
	L = iota
)

func GetNeighbours(g []string, p Point) []Point {
	n := []Point{}
	if utils.IsWithinBounds2(p.I+1, p.J, g) {
		if g[p.I+1][p.J] == g[p.I][p.J] {
			n = append(n, Point{p.I + 1, p.J})
		}
	}
	if utils.IsWithinBounds2(p.I-1, p.J, g) {
		if g[p.I-1][p.J] == g[p.I][p.J] {
			n = append(n, Point{p.I - 1, p.J})
		}
	}
	if utils.IsWithinBounds2(p.I, p.J+1, g) {
		if g[p.I][p.J+1] == g[p.I][p.J] {
			n = append(n, Point{p.I, p.J + 1})
		}
	}
	if utils.IsWithinBounds2(p.I, p.J-1, g) {
		if g[p.I][p.J-1] == g[p.I][p.J] {
			n = append(n, Point{p.I, p.J - 1})
		}
	}
	return n
}

func CalculateBorderNum(n int) int {
	return 4 - n
}

func DetermineGroupSize(g []string, s Point, v *[]Point) (int, int) {
	q := []Point{s}
	gr := []Point{}
	for {
		if len(q) == 0 {
			break
		}
		current := q[0]
		q = q[1:]
		if slices.Contains(gr, current) {
			continue
		}
		gr = append(gr, current)
		*v = append((*v), current)
		for _, n := range GetNeighbours(g, current) {
			if slices.Contains((*v), n) {
				continue
			}
			q = append(q, n)
		}
	}
	t := 0
	for _, p := range gr {
		t += CalculateBorderNum(len(GetNeighbours(g, p)))
	}
	return t * len(gr), CalculateCornerNum(gr)
}

type CornerPoint struct {
	Parent Point
	Point  Point
}

func ExpandCoordiantes(points []Point) []CornerPoint {
	new := []CornerPoint{}
	for _, p := range points {
		new = append(new, []CornerPoint{{p, Point{p.I, p.J}}, {p, Point{p.I, p.J + 1}}, {p, Point{p.I + 1, p.J}}, {p, Point{p.I + 1, p.J + 1}}}...)
	}
	return new
}

func CalculateCornerNum(group []Point) int {
	e := ExpandCoordiantes(group)
	counts := map[Point][]Point{}
	for _, p := range e {
		_, found := counts[p.Point]
		if !found {
			counts[p.Point] = []Point{}
		}
		counts[p.Point] = append(counts[p.Point], p.Parent)
	}
	t := 0
	for _, v := range counts {
		if len(v)%2 == 1 {
			t += 1
		} else if len(v) == 2 && NotNeighbours(v[0], v[1]) {
			fmt.Println(v)
			t += 2
		}
	}
	fmt.Println(t, len(group))
	return t * len(group)
}

func NotNeighbours(p1, p2 Point) bool {
	return utils.Abs(p1.I, p2.I) == 1 && utils.Abs(p1.J, p2.J) == 1
}

func Execute(filepath string) (int, int, error) {
	raw, _ := os.ReadFile(filepath)
	garden := strings.Split(string(raw), "\n")
	visited := []Point{}
	p1 := 0
	p2 := 0
	for i, row := range garden {
		for j := range row {
			p := Point{i, j}
			if slices.Contains(visited, p) {
				continue
			}
			visited = append(visited, p)
			_p1, _p2 := DetermineGroupSize(garden, p, &visited)
			p1 += _p1
			p2 += _p2
		}
	}
	return p1, p2, nil
}
