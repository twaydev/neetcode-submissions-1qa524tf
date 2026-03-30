func findMin(nums []int) int {
n := len(nums)
	if n == 1 {
		return nums[0]
	}
	res := nums[0]
	l, r := 0, n-1
	// [2,3,4,5,6,7,1] => r => mid > l > r
	// [3,4,5,6,7,1,2] => r => mid > l > r
	// [4,5,6,7,1,2,3] => r => mid > l > r
	// [5,6,7,1,2,3,4] => l => mid < r < l
	// [6,7,1,2,3,4,5] => l => mid < r < l
	// [7,1,2,3,4,5,6] => l => mid < r < l
	for l <= r {
		mid := (r + l) / 2
		nLeft, nRight := nums[l], nums[r]
		//fmt.Println(l, r, mid, "-", nums[mid], nLeft, nRight, "-", res)
		if nums[mid] <= nLeft && nums[mid] <= nRight { // move to left
			r = mid - 1
		} else if nums[mid] > nLeft && nums[mid] > nRight { // move to right
			l = mid + 1
		} else {
			if nLeft < nRight { // move to left
				r = mid - 1
			} else { // move to right
				l = mid + 1
			}
		}
		if nums[mid] < res {
			res = nums[mid]
		}
	}
	return res
}
