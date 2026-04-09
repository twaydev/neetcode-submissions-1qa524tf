
func backtrack(start int, curSum int, cur []int, target int, nums []int, res *[][]int, depth int) {
	// space := ""
	// for range depth {
	// 	space += "-"
	// }
	// fmt.Printf("%vstart: %v, curSum: %v, cur: %v\n", space, start, curSum, cur)
	if curSum == target {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	if curSum > target {
		return
	}
	depth += 1
	for i := start; i < len(nums); i++ {
		cur = append(cur, nums[i])
		backtrack(i, curSum+nums[i], cur, target, nums, res, depth)
		cur = cur[:len(cur)-1]
	}
}

func combinationSum(nums []int, target int) [][]int {
	res := [][]int{}
	backtrack(0, 0, []int{}, target, nums, &res, 0)
	return res
}