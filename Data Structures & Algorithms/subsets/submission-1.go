
func subsets(nums []int) [][]int {
// delarce result with empty subset
	result := [][]int{{}}

	for _, num := range nums {
		n := len(result)
		// fmt.Printf("\nnum=%d, result=%v", num, result)
		for i := range n {
			newSubset := make([]int, len(result[i]))
			copy(newSubset, result[i])
			newSubset = append(newSubset, num)
			// fmt.Printf("\n  %v", newSubset)
			result = append(result, newSubset)
		}
	}

	return result
}
