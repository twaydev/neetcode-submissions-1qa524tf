func subsetsWithDup(nums []int) [][]int {
// sort asc// sort asc
	sort.Ints(nums)
	result := [][]int{{}}
	start, end := 0, 0
	for j := range nums {

		start = 0

		// check duplicated
		if j > 0 && nums[j] == nums[j-1] {
			start = end
		}
		end = len(result)

		for i := start; i < end; i++ {
			newSubset := append([]int{}, result[i]...)
			newSubset = append(newSubset, nums[j])
			result = append(result, newSubset)
		}
	}
	return result
}
