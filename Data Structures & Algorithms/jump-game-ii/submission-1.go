func jump(nums []int) int {
    n := len(nums)
	// special case
	if n <= 1 {
		return 0
	}
	jumps, end, farthest := 0, 0, 0
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == end {
			end = farthest
			jumps++
			if end >= n-1 {
				break
			}
		}
	}
	return jumps
}
