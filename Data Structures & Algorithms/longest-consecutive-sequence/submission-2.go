func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 1
	hashset := make(map[int]struct{})
	// remove duplicates
	for _, n := range nums {
		hashset[n] = struct{}{}
	}

	for num := range hashset {
		// only processing if current num is the start of a sequence
		if _, hasPrev := hashset[num-1]; !hasPrev {
			length := 1
			_, hasNext := hashset[num+length]
			for hasNext {
				length++
				_, hasNext = hashset[num+length] // check next num
			}
			if length > result {
				result = length
			}
		}
	}

	return result
}