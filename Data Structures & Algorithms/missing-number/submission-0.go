func missingNumber(nums []int) int {
	h := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		h[n] = struct{}{}
	}
	//fmt.Printf("%v\n", h)
	found := 0
	for i := 0; i <= len(nums); i++ {
		if _, ok := h[i]; !ok {
			found = i
		}
	}
	return found
}