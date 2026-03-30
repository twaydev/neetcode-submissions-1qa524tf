func findDuplicate(nums []int) int {
    
	duplicated := 0
	hashmap := make(map[int]bool)
	for _, num := range nums {
		if _, ok := hashmap[num]; ok {
			duplicated = num
			break
		}
		hashmap[num] = true
	}
	return duplicated
}
