func missingNumber(nums []int) int {
	n := len(nums)
	xor := n
	for i := range n {
		xor ^= i ^ nums[i]
	}
	return xor
}