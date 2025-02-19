package day16

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Point struct {
	I int
	J int
}

type Distance struct {
	Point     Point
	Weight    float64
	Direction Direction
}

type Direction struct {
	I int
	J int
}

type PointDirection struct {
	Point     Point
	Direction Direction
}

func Execute(filepath string) (int, int, error) {
	raw, _ := os.ReadFile(filepath)
	s := strings.Split(string(raw), "\n")
	maze := make([][]string, len(s))
	start := Point{len(maze) - 2, 1}
	prev := map[Point][]Point{}
	am := map[Point]float64{}
	v := map[PointDirection]bool{}
	for i, l := range s {
		for j, r := range strings.Trim(l, "\n\r") {
			maze[i] = append(maze[i], string(r))
			am[Point{i, j}] = math.Inf(0)
			v[PointDirection{Point{i, j}, UP}] = false
			v[PointDirection{Point{i, j}, LEFT}] = false
			v[PointDirection{Point{i, j}, DOWN}] = false
			v[PointDirection{Point{i, j}, RIGHT}] = false
			prev[Point{i, j}] = []Point{}
		}
	}
	end := Point{1, len(maze[0]) - 2}
	am[start] = 0

	var mdh DistanceMinHeap
	heap.Init(&mdh)
	heap.Push(&mdh, &Distance{start, 0, RIGHT})
	for {
		if len(mdh) == 0 {
			break
		}
		current := heap.Pop(&mdh).(*Distance)
		fmt.Println(current)
		fmt.Scanln()
		v[PointDirection{current.Point, current.Direction}] = true
		for _, n := range GetNeighbours(current.Point, maze) {
			if v[PointDirection{n, current.Direction}] || slices.Contains(prev[current.Point], n) {
				continue
			}
			if n.I-current.Point.I == current.Direction.I && n.J-current.Point.J == current.Direction.J {
				heap.Push(&mdh, &Distance{n, current.Weight + 1, current.Direction})
				if am[n] > current.Weight+1 {
					am[n] = current.Weight + 1
					prev[n] = []Point{current.Point}
				}
			} else {
				heap.Push(&mdh, &Distance{n, current.Weight + 1001, Direction{n.I - current.Point.I, n.J - current.Point.J}})
				if am[n] > current.Weight+1001 {
					am[n] = current.Weight + 1001
					prev[n] = []Point{current.Point}
				}
			}
		}
	}
	unique := Part2(end, prev)
	Display(maze, unique)
	for _, v := range unique {
		fmt.Println(v, am[v])
	}
	return int(am[end]), len(unique), nil
}

func GetNeighbours(p Point, maze [][]string) []Point {
	points := []Point{}
	for _, dir := range DIRECTIONS {
		if maze[p.I+dir.I][p.J+dir.J] != "#" {
			points = append(points, Point{p.I + dir.I, p.J + dir.J})
		}
	}
	return points
}

func Display(maze [][]string, unique []Point) {
	for i, r := range maze {
		for j, c := range r {
			if slices.Contains(unique, Point{i, j}) {
				fmt.Print("\033[31mO\033[0m")
			} else {
				fmt.Print(c)
			}

		}
		fmt.Println()
	}
}

func Part2(end Point, prev map[Point][]Point) []Point {
	fmt.Println(prev)
	q := Queue{end}
	unique := []Point{}
	for {
		if len(q) == 0 {
			break
		}
		c := q.Pop()
		if c.I == 0 && c.J == 0 {
			continue
		}
		if !slices.Contains(unique, c) {
			unique = append(unique, c)
		}
		q = append(q, prev[c]...)
		prev[c] = []Point{}
	}
	fmt.Println(unique)
	return unique
}
