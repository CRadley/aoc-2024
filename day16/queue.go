package day16

type Queue []Point

func (heap *Queue) Push(x Point) {
	*heap = append(*heap, x)
}

func (heap *Queue) Pop() Point {
	temp := *heap
	*heap = temp[1:]
	return temp[0]
}
