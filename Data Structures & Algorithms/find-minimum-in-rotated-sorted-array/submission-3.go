func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (r + l) / 2
		if nums[m] < nums[r] { // move to left including m
			r = m
		} else { // move to right excluding m
			l = m + 1
		}
	}
	return nums[l]
}
