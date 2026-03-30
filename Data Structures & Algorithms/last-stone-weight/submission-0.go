
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // max-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	// special case
	if len(stones) == 1 {
		return stones[0]
	}

	// build heap
	h := &IntHeap{}
	*h = append(*h, stones...)
	heap.Init(h)

	// now we have max-heap
	// get max weight stone
	for len(*h) > 0 {
		stone1 := heap.Pop(h).(int)
		if len(*h) == 0 {
			return stone1
		}
		stone2 := heap.Pop(h).(int)
		if stone1 != stone2 {
			heap.Push(h, stone1-stone2)
		}
	}

	return 0
}