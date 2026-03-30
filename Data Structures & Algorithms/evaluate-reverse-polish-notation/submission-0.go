func evalRPN(tokens []string) int {
	result := []int{}
	for _, char := range tokens {
		if char == "+" || char == "-" || char == "*" || char == "/" {
			// get 2 elements from result
			num2 := result[len(result)-1]
			num1 := result[len(result)-2]
			result = result[:len(result)-2]
			switch char {
				case "+":
				result = append(result, num1+num2)
				case "-":
				result = append(result, num1-num2)
				case "*":
				result = append(result, num1*num2)
				case "/":
				result = append(result, num1/num2)
			}
		} else {
			num, _ := strconv.Atoi(char)
			result = append(result, num)
		}
	}
	return result[0]
}
