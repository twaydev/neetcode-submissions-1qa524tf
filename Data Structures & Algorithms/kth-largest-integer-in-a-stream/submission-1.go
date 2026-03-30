

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // min heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(i any)        { *h = append(*h, i.(int)) }
func (h *IntHeap) Pop() any {
	old := *h         // copy values
	n := len(old)     // get len of old
	x := old[n-1]     // get root to return. Because the parent pop func has swap begin - end el
	*h = old[0 : n-1] // remove root from heap and update heap
	return x
}

type KthLargest struct {
	heap *IntHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	var h IntHeap
	if len(nums) > 0 {
		h = append(h, nums...)
	}
	heap.Init(&h)
	for len(h) > k {
		heap.Pop(&h)
	}
	return KthLargest{
		heap: &h,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	// fmt.Printf("heap: %v\n", this.heap)
	heap.Push(this.heap, val)
	if len(*this.heap) > this.k {
		heap.Pop(this.heap) // remove the smallest
	}
	return (*this.heap)[0]
}
