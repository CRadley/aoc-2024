package day16

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	I uint
	J uint
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
	Point  Point
	Weight uint
	Facing Point
}

type VisitedDirection struct {
	Point  Point
	Facing Point
}

func Execute(filepath string) (int, int, error) {
	raw, _ := os.ReadFile(filepath)
	s := strings.Split(string(raw), "\n")
	maze := make([][]string, len(s))
	start := Point{uint(len(maze) - 2), 1}
	prev := map[Point][]Point{}
	am := map[Point]float64{}
	v := map[Point]bool{}
	for i, l := range s {
		for j, r := range l {
			maze[i] = append(maze[i], string(r))
			if string(r) != "#" {
				am[Point{uint(i), uint(j)}] = math.Inf(0)
				v[Point{uint(i), uint(j)}] = false
				prev[Point{uint(i), uint(j)}] = []Point{}
			}
		}
	}
	end := Point{1, uint(len(maze[0]) - 3)}
	am[start] = 0

	var mdh DistanceMinHeap
	heap.Init(&mdh)
	heap.Push(&mdh, &Distance{start, 0, Point{0, 1}})
	for {
		if len(mdh) == 0 {
			break
		}

		current := heap.Pop(&mdh).(*Distance)
		v[current.Point] = true
		for _, n := range GetNeighbours(current.Point, maze) {
			if v[n] {
				continue
			}
			if n.I-current.Point.I == current.Facing.I && n.J-current.Point.J == current.Facing.J {
				heap.Push(&mdh, &Distance{n, current.Weight + 1, Point{n.I - current.Point.I, n.J - current.Point.J}})
				if am[n] >= float64(current.Weight)+1 {
					if am[n] > float64(current.Weight)+1 {
						prev[n] = []Point{current.Point}
					} else {
						prev[n] = append(prev[n], current.Point)
					}
					am[n] = float64(current.Weight) + 1
				}
			} else {
				heap.Push(&mdh, &Distance{n, current.Weight + 1001, Point{n.I - current.Point.I, n.J - current.Point.J}})
				if am[n] >= float64(current.Weight)+1001 {

					if am[n] > float64(current.Weight)+1001 {
						prev[n] = []Point{current.Point}
					} else {
						prev[n] = append(prev[n], current.Point)
					}
					am[n] = float64(current.Weight) + 1001
				}
			}
		}
	}
	fmt.Println(prev)
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
