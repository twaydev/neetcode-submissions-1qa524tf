
type Point struct {
	distance int
	point    []int
}

type ClosetHeap []Point

func (h ClosetHeap) Len() int           { return len(h) }
func (h ClosetHeap) Less(i, j int) bool { return h[i].distance > h[j].distance } // max heap
func (h ClosetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ClosetHeap) Push(i any)        { *h = append(*h, i.(Point)) }
func (h *ClosetHeap) Pop() any {
	old := *h      // copy values
	n := len(old)  // get len of old
	x := old[n-1]  // get root to return. Because the parent pop func has swap begin - end el
	*h = old[:n-1] // remove root from heap and update heap
	return x
}

func kClosest(points [][]int, k int) [][]int {

	h := &ClosetHeap{}
	heap.Init(h)
	for _, p := range points {
		// distance := math.Sqrt(float64(p[0]*p[0] + p[1]*p[1]))
		// fmt.Println(distance, p)
		heap.Push(h, Point{
			distance: p[0]*p[0] + p[1]*p[1],
			point:    p,
		})
		// remove farest
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([][]int, h.Len())
	for i := range res {
		res[i] = heap.Pop(h).(Point).point
	}

	return res
}