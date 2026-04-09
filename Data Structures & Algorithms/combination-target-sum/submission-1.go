
func combinationSum(nums []int, target int) [][]int {
	res := [][]int{}
	// backtrack(0, 0, []int{}, target, nums, &res, 0)
	var backtrack func(start int, curSum int, cur []int)
	backtrack = func(start int, curSum int, cur []int) {
		if curSum == target {
			res = append(res, append([]int{}, cur...))
			return
		}
		if curSum > target {
			return
		}
		for i := start; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrack(i, curSum+nums[i], cur)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0, 0, []int{})
	return res
}