package utils

func Abs(x, y int) int {
	value := x - y
	if value < 0 {
		return -value
	}
	return value
}
