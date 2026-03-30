func search(nums []int, target int) int {
mid := len(nums) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		for i := mid + 1; i < len(nums); i++ {
			if nums[i] == target {
				return i
			}
		}
	} else {
		for i := 0; i <= mid-1; i++ {
			if nums[i] == target {
				return i
			}
		}
	}

	return -1
}
