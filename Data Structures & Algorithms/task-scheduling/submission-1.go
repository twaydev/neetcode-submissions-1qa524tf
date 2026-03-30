
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

type DelayTask struct {
	remain int
	time   int
}

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return 1
	}
	freqAZ := [26]int{}
	for i := range tasks {
		freqAZ[tasks[i]-'A']++
	}
	// fmt.Printf("%v\n", freqAZ)

	// declare max-heap with freqAZ
	h := &IntHeap{}
	heap.Init(h)
	for i := range freqAZ {
		if freqAZ[i] > 0 {
			heap.Push(h, freqAZ[i])
		}
	}

	// fmt.Printf("%v\n", h)
	delayQueue := []DelayTask{}
	i := 0
	for len(*h) > 0 || len(delayQueue) > 0 {
		i++
		if len(*h) > 0 {
			// get task with max feq
			task := heap.Pop(h).(int)
			task--
			if task > 0 {
				delayQueue = append(delayQueue, DelayTask{remain: task, time: i + n})
			}
		}
		if len(delayQueue) > 0 && delayQueue[0].time == i {
			heap.Push(h, delayQueue[0].remain)
			delayQueue = delayQueue[1:]
		}
	}
	return i
}