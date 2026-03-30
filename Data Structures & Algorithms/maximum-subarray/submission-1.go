func maxSubArray(nums []int) int {
    // special cases
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// declare maxVal to return
	maxVal := nums[0]
	// temp maxVal to update while itegrating nums
	tmpMaxVal := nums[0]
	// use greedy to compare tmpMaxVal+nums[i] vs nums[i]
	for i := 1; i < len(nums); i++ {
		if tmpMaxVal+nums[i] >= nums[i] {
			tmpMaxVal += nums[i]
		} else {
			tmpMaxVal = nums[i]
		}
		if tmpMaxVal > maxVal {
			maxVal = tmpMaxVal
		}
		// fmt.Printf("tmpMaxVal: %v, maxVal: %v\n", tmpMaxVal, maxVal)
	}
	return maxVal
}
