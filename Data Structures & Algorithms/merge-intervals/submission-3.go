func merge(intervals [][]int) [][]int {
    n := len(intervals)
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return intervals
	}
	// sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	i := 0
	for i <= n-1 {
		// so sánh intervals[i] vs res[len(res)-1]
		if len(res) == 0 {
			res = append(res, intervals[i])
		} else if res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else { // merge
			res[len(res)-1][0] = min(res[len(res)-1][0], intervals[i][0])
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		}
		i++
	}
	return res
}
