package main

import (
	"aoc2024/day01"
	"aoc2024/day02"
	"aoc2024/day03"
	"aoc2024/day04"
	"aoc2024/day05"
	"fmt"
	"os"
	"runtime/debug"
)

func executeDay(day string) (int, int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Something went wrong... \n\n%s\n\n%s", r, debug.Stack())
			os.Exit(-1)
		}
	}()
	filepath := fmt.Sprintf("input/%s", day)
	switch day {
	case "1":
		return day01.Execute(filepath)
	case "2":
		return day02.Execute(filepath)
	case "3":
		return day03.Execute(filepath)
	case "4":
		return day04.Execute(filepath)
	case "5":
		return day05.Execute(filepath)
	}

	return -1, -1, fmt.Errorf("day '%s' not found", day)
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		os.Exit(1)
	}
	p1, p2, err := executeDay(args[0])
	fmt.Printf("---Part 1---\n%d\n---Part 2---\n%d\n---ERROR---\n%s\n", p1, p2, err)
}
