func threeSum(nums []int) [][]int {
	var result [][]int
	// sort nums ascending
	// slices.SortFunc(nums, func(a, b int) int {
	// 	return a - b
	// })
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i, n := range nums {
		// check exist, if true, move to next
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// use two pointers for lookup the pair of j and k
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right]
			if -sum == n {
				item := []int{n, nums[left], nums[right]}
				result = append(result, item)
				// skip duplicate values for the left pointer
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				// skip duplicate values for the left pointer
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// continue processing
				right--
				left++
			} else if -sum < n {
				right--
			} else {
				left++
			}
		}
	}
	if len(result) == 0 {
		return [][]int{}
	}
	return result
}