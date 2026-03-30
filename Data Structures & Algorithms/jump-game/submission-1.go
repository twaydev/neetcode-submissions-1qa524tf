func canJump(nums []int) bool {
    n := len(nums)

	targetIndex := n - 1
	for i := n - 2; i >= 0; i-- {
		if nums[i]+i >= targetIndex {
			targetIndex = i
		}
	}
	return targetIndex == 0

}
