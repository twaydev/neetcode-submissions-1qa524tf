func partitionLabels(s string) []int {
   // last occurrence
	lastOccurence := [26]int{}
	for i, n := range s {
		lastOccurence[n-'a'] = i
	}
	// delcare partitions
	partitions := []int{}
	// logic to create partitions
	currentStart, currEnd := 0, 0
	for i, n := range s {
		currEnd = max(currEnd, lastOccurence[n-'a'])
		if i == currEnd {
			partitions = append(partitions, currEnd-currentStart+1)
			currentStart = i + 1
		}
	}
	return partitions 
}
