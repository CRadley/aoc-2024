package utils

func Abs(x, y int) int {
	value := x - y
	if value < 0 {
		return -value
	}
	return value
}

func IsWithinBounds(i, j int, lines [][]string) bool {
	return (i >= 0 && i < len(lines)) && (j >= 0 && j < len(lines[0]))
}
