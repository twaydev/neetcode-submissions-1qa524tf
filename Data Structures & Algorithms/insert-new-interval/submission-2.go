func insert(intervals [][]int, newInterval []int) [][]int {
        res := [][]int{}
    i, n := 0, len(intervals)

    // 1) Add all intervals that end before newInterval starts
    for i < n && intervals[i][1] < newInterval[0] {
        res = append(res, intervals[i])
        i++
    }

    // 2) Merge all intervals that overlap with newInterval
    for i < n && intervals[i][0] <= newInterval[1] {
        if intervals[i][0] < newInterval[0] {
            newInterval[0] = intervals[i][0]
        }
        if intervals[i][1] > newInterval[1] {
            newInterval[1] = intervals[i][1]
        }
        i++
    }
    // push merged newInterval
    res = append(res, newInterval)

    // 3) Add all remaining intervals (to the right)
    res = append(res, intervals[i:]...)

    return res
}
