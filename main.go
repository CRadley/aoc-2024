package main

import (
	"aoc2024/day01"
	"aoc2024/day02"
	"aoc2024/day03"
	"aoc2024/day04"
	"aoc2024/day05"
	"aoc2024/day06"
	"aoc2024/day07"
	"aoc2024/day08"
	"aoc2024/day09"
	"aoc2024/day10"
	"aoc2024/day11"
	"aoc2024/day12"
	"aoc2024/day13"
	"aoc2024/day14"
	"fmt"
	"os"
	"runtime/debug"
	"time"
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
	case "6":
		return day06.Execute(filepath)
	case "7":
		return day07.Execute(filepath)
	case "8":
		return day08.Execute(filepath)
	case "9":
		return day09.Execute(filepath)
	case "10":
		return day10.Execute(filepath)
	case "11":
		return day11.Execute(filepath)
	case "12":
		return day12.Execute(filepath)
	case "13":
		return day13.Execute(filepath)
	case "14":
		return day14.Execute(filepath)
	}
	return -1, -1, fmt.Errorf("day '%s' not found", day)
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		os.Exit(1)
	}
	start := time.Now()
	p1, p2, err := executeDay(args[0])
	fmt.Printf("---Part 1---\n%d\n---Part 2---\n%d\n---ERROR---\n%s\n---Time Taken---\n%s\n", p1, p2, err, time.Since(start))
}
