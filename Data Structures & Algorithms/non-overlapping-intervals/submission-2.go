func eraseOverlapIntervals(intervals [][]int) int {
   res := 0
	if len(intervals) == 0 || len(intervals) == 1 {
		return res
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	prevEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= prevEnd {
			prevEnd = intervals[i][1]
		} else {
			res++
		}
	}
	return res
}
