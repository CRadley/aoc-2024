package day13

import (
	"os"
	"regexp"
	"sort"
	"strconv"
)

type ClawMachine struct {
	AX int
	AY int
	BX int
	BY int
	PX int
	PY int
}

func Execute(filepath string) (int, int, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return -1, -1, err
	}
	data := string(bytes)
	p, _ := regexp.Compile(`(\d+)`)
	matches := p.FindAllString(data, -1)
	machines := []ClawMachine{}
	machines2 := []ClawMachine{}
	ints := []int{}
	for _, m := range matches {
		val, _ := strconv.Atoi(m)
		ints = append(ints, val)
	}
	for {
		if len(ints) <= 0 {
			break
		}
		machines = append(machines, ClawMachine{ints[0], ints[1], ints[2], ints[3], ints[4], ints[5]})
		machines2 = append(machines2, ClawMachine{ints[0], ints[1], ints[2], ints[3], ints[4] + 10000000000000, ints[5] + 10000000000000})
		ints = ints[6:]
	}
	p1 := 0
	p2 := 0
	for _, machine := range machines {
		tracker := []int{}
		for i := range 101 {
			if (machine.PX-(machine.AX*i))%machine.BX == 0 && (machine.PY-(machine.AY*i))%machine.BY == 0 {
				if (machine.PX-(machine.AX*i))/machine.BX == (machine.PY-(machine.AY*i))/machine.BY {
					tracker = append(tracker, (i*3)+(machine.PY-(machine.AY*i))/machine.BY)
				}
				continue
			}
		}
		if len(tracker) > 0 {
			sort.Ints(tracker)
			p1 += tracker[0]
		}
	}
	for _, machine := range machines2 {
		PX2 := machine.PX * machine.AY
		BX2 := machine.BX * machine.AY
		PY2 := machine.PY * machine.AX
		BY2 := machine.BY * machine.AX
		D := PX2 - PY2
		D2 := BX2 - BY2
		B := D / D2
		A := (machine.PX - (machine.BX * B)) / machine.AX
		total := A*3 + B
		if total > 0 && A >= 0 && B >= 0 && A*machine.AX+B*machine.BX == machine.PX && A*machine.AY+B*machine.BY == machine.PY {
			p2 += total
		}
	}
	return p1, p2, nil
}
