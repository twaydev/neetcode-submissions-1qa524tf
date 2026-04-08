func subsetsWithDup(nums []int) [][]int {
// sort asc// sort asc
	sort.Ints(nums)
	result := [][]int{{}}
	start, end := 0, 0
	for j, num := range nums {
		n := len(result)

		start = 0

		// check duplicated
		if j > 0 && num == nums[j-1] {
			start = end
		}
		end = n

		for i := start; i < end; i++ {
			newSubset := make([]int, len(result[i]))
			copy(newSubset, result[i])
			newSubset = append(newSubset, num)
			result = append(result, newSubset)
		}
	}
	return result
}
