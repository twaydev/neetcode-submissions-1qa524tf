
func get2Elements(result []int) (num1 int, num2 int, res []int) {
	return result[len(result)-2],result[len(result)-1], result[:len(result)-2]
}
func evalRPN(tokens []string) int {
	result := []int{}
	for _, char := range tokens {
		switch char {
		case "+":
		num1, num2, newResult := get2Elements(result)
		result = append(newResult, num1+num2)
		case "-":
		num1, num2, newResult := get2Elements(result)
		result = append(newResult, num1-num2)
		case "*":
		num1, num2, newResult := get2Elements(result)
		result = append(newResult, num1*num2)
		case "/":
		num1, num2, newResult := get2Elements(result)
		result = append(newResult, num1/num2)
		default:
		num, _ := strconv.Atoi(char)
		result = append(result, num)
		}
	}
	return result[0]
}