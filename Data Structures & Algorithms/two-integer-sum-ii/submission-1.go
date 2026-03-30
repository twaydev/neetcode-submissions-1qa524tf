
func twoSum(numbers []int, target int) []int {

	// declare 2 variables from left and right
	var i = 0
	j := len(numbers) - 1

	// loop to look up the 2 indices
	for i < j {
		for i < j {

			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1} // return 1-indexed
			}
			j--

		}
		i++
		j = len(numbers) - 1

	}

	return nil
}