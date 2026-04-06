
func backtrack(nums []int, start int, path []int, result *[][]int) {

	fmt.Printf("%v - %v\n", start, path)
	// thêm tập con hiện tại vào result
	subset := make([]int, len(path))
	copy(subset, path)
	*result = append(*result, subset)

	// thử từng phần tử từ vị trí start
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrack(nums, i+1, path, result)
		path = path[:len(path)-1]
	}
}

func subsets(nums []int) [][]int {
	result := [][]int{}
	backtrack(nums, 0, []int{}, &result)
	return result
}
