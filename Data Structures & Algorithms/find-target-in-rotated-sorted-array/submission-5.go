func search(nums []int, target int) int {
	
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		// 3, 5, 6, 0, 1, 2
		// fmt.Println(mid, nums[mid], target, left, right)
		if nums[mid] == target {
			return mid
		}
		// Nửa phải sorted (nums[mid] <= nums[right])
		if nums[mid] <= nums[right] {
			if target > nums[right] || target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // Nửa trái sorted (nums[left] <= nums[mid])
			if target < nums[left] || target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
