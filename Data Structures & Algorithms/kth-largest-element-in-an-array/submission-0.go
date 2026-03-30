
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(i any)        { *h = append(*h, i.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	l := len(*h)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func findKthLargest(nums []int, k int) int {

	if len(nums) == 0 {
		return -1
	}
	var h IntHeap
	h = append(h, nums...)
	heap.Init(&h)
	count := 1 // 1-indexed
	for count < k {
		x := heap.Pop(&h)
		fmt.Println(x)
		count++
	}

	return h[0]
}