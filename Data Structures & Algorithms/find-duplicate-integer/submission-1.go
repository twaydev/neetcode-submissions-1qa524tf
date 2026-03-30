func findDuplicate(nums []int) int {
    fast, slow := 0, 0

	// find meeting point
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if fast == slow {
			break // found
		}
	}

	// find entry point of cycle
	slow2 := 0
	for {
		slow2 = nums[slow2]
		slow = nums[slow]
		if slow == slow2 {
			return slow
		}
	}
	return 0
}
