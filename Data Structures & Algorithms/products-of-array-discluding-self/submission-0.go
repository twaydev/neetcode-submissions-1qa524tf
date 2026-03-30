func productExceptSelf(nums []int) []int {

	if len(nums) == 2 {
		return []int{nums[1], nums[0]}
	}
	// initialize result array [0, 0, 0 ...] with len from nums
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		j := 0
		result[i] = 1
		for j < len(nums) {
			if i != j {
				result[i] *= nums[j]
			}
			j++
		}

	}
	return result
}
