package day10

import (
	"aoc2024/utils"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	I int
	J int
}

func FindStartingPoints(m [][]int) []Point {
	points := []Point{}
	for i, row := range m {
		for j, v := range row {
			if v == 0 {
				points = append(points, Point{i, j})
			}
		}
	}
	return points
}

func DetermineScore(m [][]int, start Point) (int, int) {
	finalPoints := []Point{}
	finalPoints2 := []Point{}
	visited := []Point{}
	queue := []Point{start}
	for {
		if len(queue) == 0 {
			break
		}
		current := queue[0]
		queue = queue[1:]
		visited = append(visited, current)
		if m[current.I][current.J] == 9 {
			if !slices.Contains(finalPoints, current) {
				finalPoints = append(finalPoints, current)
			}
			finalPoints2 = append(finalPoints2, current)
		}
		for _, n := range GetNeighbours(m, current) {
			if slices.Contains(visited, n) {
				continue
			}
			queue = append(queue, n)
		}
	}
	return len(finalPoints), len(finalPoints2)
}

func GetNeighbours(m [][]int, current Point) []Point {
	n := []Point{}
	if utils.IsWithinBounds3(current.I+1, current.J, m) {
		if m[current.I+1][current.J]-m[current.I][current.J] == 1 {
			n = append(n, Point{current.I + 1, current.J})
		}
	}
	if utils.IsWithinBounds3(current.I-1, current.J, m) {
		if m[current.I-1][current.J]-m[current.I][current.J] == 1 {
			n = append(n, Point{current.I - 1, current.J})
		}
	}
	if utils.IsWithinBounds3(current.I, current.J+1, m) {
		if m[current.I][current.J+1]-m[current.I][current.J] == 1 {
			n = append(n, Point{current.I, current.J + 1})
		}
	}
	if utils.IsWithinBounds3(current.I, current.J-1, m) {
		if m[current.I][current.J-1]-m[current.I][current.J] == 1 {
			n = append(n, Point{current.I, current.J - 1})
		}
	}
	return n
}

func Execute(filepath string) (int, int, error) {
	m := [][]int{}
	raw, _ := os.ReadFile(filepath)
	data := strings.Split(string(raw), "\n")
	for _, row := range data {
		values := []int{}
		for _, val := range row {
			x, _ := strconv.Atoi(string(val))
			values = append(values, x)
		}
		m = append(m, values)
	}
	p1 := 0
	p2 := 0
	startingPoints := FindStartingPoints(m)
	for _, s := range startingPoints {
		l, r := DetermineScore(m, s)
		p1 += l
		p2 += r
	}
	return p1, p2, nil
}
