package day16

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
