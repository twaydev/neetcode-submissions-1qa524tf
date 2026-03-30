func dailyTemperatures(temperatures []int) []int {
// special scheck
	n := len(temperatures)
	if n <= 1 {
		return []int{0}
	}
	// declare result with n elements and default value as 0
	result := make([]int, n)

	// declare monotonic decreasing stack
	stack := []int{} // lưu chỉ số
	for i := 0; i <= n-1; i++ {
		//fmt.Printf("i=%d, result=%v, stack=%v\n", i, result, stack)
		// we run while loop to look for next greater value
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = i - idx
			//fmt.Printf("     result=%v\n", result)
		}
		// push for each step
		stack = append(stack, i)
	}
	return result
}
