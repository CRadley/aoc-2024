package day11

import (
	"os"
	"strconv"
	"strings"
)

type State struct {
	E int
	C int
}

func GetStates(m map[int]int) []State {
	s := []State{}
	for k, v := range m {
		s = append(s, State{k, v})
	}
	return s
}

func Execute(filepath string) (int, int, error) {
	m := map[int]int{}
	raw, _ := os.ReadFile(filepath)
	for _, v := range strings.Split(string(raw), " ") {
		x, _ := strconv.Atoi(string(v))
		m[x] = 1
	}
	counts := map[int]int{}
	for i := range 75 {
		states := GetStates(m)
		for _, state := range states {
			if state.E == 0 {
				m[0] -= state.C
				_, found := m[1]
				if !found {
					m[1] = 0
				}
				m[1] += state.C
			} else if len(strconv.Itoa(state.E))%2 == 0 {
				m[state.E] -= state.C
				s := strconv.Itoa(state.E)
				s0 := s[:len(s)/2]
				s1 := s[len(s)/2:]
				s0i, _ := strconv.Atoi(s0)
				s1i, _ := strconv.Atoi(s1)
				_, found := m[s0i]
				if !found {
					m[s0i] = 0
				}
				m[s0i] += state.C
				_, found = m[s1i]
				if !found {
					m[s1i] = 0
				}
				m[s1i] += state.C
			} else {
				m[state.E] -= state.C
				_, found := m[state.E*2024]
				if !found {
					m[state.E*2024] = 0
				}
				m[state.E*2024] += state.C
			}
		}
		for _, v := range m {
			counts[i+1] += v
		}
	}
	return counts[25], counts[75], nil
}
