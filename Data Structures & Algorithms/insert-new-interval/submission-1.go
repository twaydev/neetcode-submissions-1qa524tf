func insert(intervals [][]int, newInterval []int) [][]int {
    // special cases
	if len(intervals) == 0 {
		if len(newInterval) == 0 {
			return [][]int{}
		}
		return [][]int{newInterval}
	}

	var res [][]int
	for i, node := range intervals {
		if newInterval[1] < node[0] {
			res = append(res, newInterval)
			return append(res, intervals[i:]...)
		} else if newInterval[0] > node[1] {
			res = append(res, node)
		} else {
			newInterval[0] = min(newInterval[0], node[0])
			newInterval[1] = max(newInterval[1], node[1])
		}
	}
	res = append(res, newInterval)
	return res
}
