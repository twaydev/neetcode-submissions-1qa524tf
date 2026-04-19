func hammingWeight(n int) int {
	res := 0
	for i := range 32 {
		if 1<<i&n != 0 {
			res++
		}
	}
	return res
}