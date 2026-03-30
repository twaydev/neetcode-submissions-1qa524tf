func twoSum(nums []int, target int) []int {
    
	// create hashtable to store indices
	hashtable := make(map[int]int)
	for i, n := range nums {
		// calculate compare value as j value
		jValue := target - n

		// lookup index of i from j value
		if j, ok := hashtable[jValue]; ok {
			return []int{j, i}
		}

		// set values to hash table
		hashtable[n] = i
	}
	
	return []int{}
}
