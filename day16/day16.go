package day16

import (
	"container/heap"
	"math"
	"os"
	"slices"
	"strings"
)

type Point struct {
	I int
	J int
}

type DistanceMinHeap []*Distance

func (heap DistanceMinHeap) Len() int {
	return len(heap)
}

func (heap DistanceMinHeap) Less(i, j int) bool {
	return heap[i].Weight < heap[j].Weight
}

func (heap DistanceMinHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *DistanceMinHeap) Push(x any) {
	*heap = append(*heap, x.(*Distance))
}

func (heap *DistanceMinHeap) Pop() any {
	temp := *heap
	l := len(temp)
	*heap = temp[:l-1]
	return temp[l-1]
}

func (heap *DistanceMinHeap) Peek() *Distance {
	return (*heap)[len(*heap)-1]
}

type Distance struct {
	Point     Point
	Weight    uint
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

type Queue []Point

func (heap *Queue) Push(x Point) {
	*heap = append(*heap, x)
}

func (heap *Queue) Pop() Point {
	temp := *heap
	*heap = temp[1:]
	return temp[0]
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
			v[PointDirection{Point{i, j}, Direction{0, 1}}] = false
			v[PointDirection{Point{i, j}, Direction{0, -1}}] = false
			v[PointDirection{Point{i, j}, Direction{-1, 0}}] = false
			v[PointDirection{Point{i, j}, Direction{1, 0}}] = false
			prev[Point{i, j}] = []Point{}
		}
	}
	end := Point{1, len(maze[0]) - 2}
	am[start] = 0

	var mdh DistanceMinHeap
	heap.Init(&mdh)
	heap.Push(&mdh, &Distance{start, 0, Direction{0, 1}})
	for {
		if len(mdh) == 0 {
			break
		}
		current := heap.Pop(&mdh).(*Distance)
		v[PointDirection{current.Point, current.Direction}] = true
		for _, n := range GetNeighbours(current.Point, maze) {
			if v[PointDirection{n, current.Direction}] || slices.Contains(prev[current.Point], n) {
				continue
			}
			if n.I-current.Point.I == current.Direction.I && n.J-current.Point.J == current.Direction.J {
				heap.Push(&mdh, &Distance{n, current.Weight + 1, current.Direction})
				if am[n] == float64(current.Weight)+1 {
					prev[n] = append(prev[n], current.Point)
				} else if am[n] > float64(current.Weight)+1 {
					am[n] = float64(current.Weight) + 1
					prev[n] = []Point{current.Point}
				}
			} else {
				heap.Push(&mdh, &Distance{n, current.Weight + 1001, Direction{n.I - current.Point.I, n.J - current.Point.J}})
				if am[n] == float64(current.Weight)+1001 {
					prev[n] = append(prev[n], current.Point)
				} else if am[n] > float64(current.Weight)+1001 {
					am[n] = float64(current.Weight) + 1001
					prev[n] = []Point{current.Point}
				}
			}
		}
	}

	return int(am[end]), -1, nil
}

func GetNeighbours(p Point, maze [][]string) []Point {
	points := []Point{}
	if maze[p.I-1][p.J] != "#" {
		points = append(points, Point{p.I - 1, p.J})
	}
	if maze[p.I+1][p.J] != "#" {
		points = append(points, Point{p.I + 1, p.J})
	}
	if maze[p.I][p.J-1] != "#" {
		points = append(points, Point{p.I, p.J - 1})
	}
	if maze[p.I][p.J+1] != "#" {
		points = append(points, Point{p.I, p.J + 1})
	}
	return points
}
