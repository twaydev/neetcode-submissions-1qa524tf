
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	var backtrack func(start, sum int, path []int)
	backtrack = func(start, currSum int, currPath []int) {
		// fmt.Printf("start: %v, currSum: %v, currPath: %v\n", start, currSum, currPath)
		if currSum == target {
			res = append(res, append([]int{}, currPath...))
			return
		}
		if currSum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			if currSum+candidates[i] > target {
				break
			}
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			currPath = append(currPath, candidates[i])
			backtrack(i+1, currSum+candidates[i], currPath)
			currPath = currPath[:len(currPath)-1]
		}
	}
	backtrack(0, 0, []int{})
	return res
}
