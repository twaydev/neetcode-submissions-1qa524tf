
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i>>1] + i&1
		// fmt.Printf("%v: %b i>>1: %v res: %v\n", i, i, i>>1, res)
	}
	return res
}