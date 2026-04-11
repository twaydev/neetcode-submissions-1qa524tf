func permute(nums []int) [][]int {
	res := [][]int{}
	var backtrack func(found []bool, path []int, depth int)
	backtrack = func(found []bool, path []int, depth int) {
		// fmt.Printf("depth: %v, found: %v, path: %v\n", depth, found, path)
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := range nums {
			if found[i] {
				continue
			}
			found[i] = true
			path = append(path, nums[i])
			backtrack(found, path, depth+1)
			found[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack(make([]bool, len(nums)), []int{}, 0)

	return res
}