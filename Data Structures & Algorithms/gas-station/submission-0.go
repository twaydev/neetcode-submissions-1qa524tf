
func sum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
func canCompleteCircuit(gas []int, cost []int) int {
	// check special case when total gas does not cover cost
	if sum(gas) < sum(cost) {
		return -1
	}
	// greedy
	currentTank := 0
	res := 0
	for i := range gas {
		currentTank += gas[i] - cost[i]
		if currentTank < 0 { // không đủ gas chạy tới i+1
			res = i + 1
			currentTank = 0
		}
	}
	return res
}