func topKFrequent(nums []int, k int) []int {

	// initialize hashmap to count frequent
	frequencies := make(map[int]int)
	for _, num := range nums {
		// count num and increase the value to hashmap
		frequencies[num]++
	}

	// initialize empty topFreq quit len and cap equal to len(nums) + 1
	// we must +1 because frequency can start from 0 to len(nums)
	topFrequencies := make([][]int, len(nums)+1)
	for num, freq := range frequencies {
		topFrequencies[freq] = append(topFrequencies[freq], num)
	}

	// retrieve k elements from topFrequencies
	fmt.Println(topFrequencies)
	result := []int{}
	for freq := len(topFrequencies) - 1; freq >= 0; freq-- {
		for _, num := range topFrequencies[freq] {
			result = append(result, num)
			if len(result) == k {
				// sort ascending
				sort.Slice(result, func(i, j int) bool {
					return result[i] < result[j]
				})
				return result
			}
		}
	}

	return result
}
