package day14

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	X  int
	Y  int
	DX int
	DY int
}

func (robot *Robot) Move(steps, width, height int) {
	robot.X = Mod(robot.X+(steps*robot.DX), width)
	robot.Y = Mod(robot.Y+(steps*robot.DY), height)
}

func Mod(a, b int) int {
	return (a%b + b) % b
}

func ContainsTree(c [][]string) bool {
	for _, v := range c {
		b := strings.Join(v, "")
		if strings.Contains(b, "###########") {
			return true
		}
	}
	return false
}

func GenerateEmptyArrays(w, h int) [][]string {
	arrays := make([][]string, h)
	for j := range arrays {
		arrays[j] = make([]string, w)
		for k := range w {
			arrays[j][k] = "."
		}
	}
	return arrays
}

func Execute(filepath string) (int, int, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return -1, -1, err
	}
	data := string(bytes)
	p, _ := regexp.Compile(`([\d-]+)`)
	matches := p.FindAllString(data, -1)
	robots := []*Robot{}
	ints := []int{}
	for _, m := range matches {
		val, _ := strconv.Atoi(m)
		ints = append(ints, val)
	}
	for {
		if len(ints) <= 0 {
			break
		}
		robots = append(robots, &Robot{ints[0], ints[1], ints[2], ints[3]})
		ints = ints[4:]
	}
	// Top Left, Bottom Left, Top Right, Bottom Right
	q1, q2, q3, q4 := 0, 0, 0, 0
	WIDTH, HEIGHT := 101, 103
	mw := (WIDTH - 1) / 2
	mh := (HEIGHT - 1) / 2
	stepCount := 0
	for {
		stepCount += 1
		arrays := GenerateEmptyArrays(WIDTH, HEIGHT)
		for _, robot := range robots {
			robot.Move(1, WIDTH, HEIGHT)
			// Part one [100 steps]
			if stepCount == 100 {
				if robot.X < mw && robot.Y < mh {
					q1 += 1
				} else if robot.X < mw && robot.Y > mh {
					q2 += 1
				} else if robot.X > mw && robot.Y < mh {
					q3 += 1
				} else if robot.X > mw && robot.Y > mh {
					q4 += 1
				}
			}
			arrays[robot.Y][robot.X] = "#"
		}
		if ContainsTree(arrays) {
			break
		}
	}
	return q1 * q2 * q3 * q4, stepCount, nil
}
