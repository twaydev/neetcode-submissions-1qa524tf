func isValid(s string) bool {
    tmp := make(map[rune]rune)
	tmp[')'] = '('
	tmp['}'] = '{'
	tmp[']'] = '['

	rString := []rune(s)
	stack := []rune{} // FILO
	// fmt.Println("rString", rString)
	for i := 0; i <= len(rString)-1; i++ {
		// fmt.Printf("%v, %v\n", tmp[rString[i]], stack)
		// this is closing bracket, look up for opening bracket
		if _, exist := tmp[rString[i]]; exist {
			if len(stack) > 0 && stack[len(stack)-1] == tmp[rString[i]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else { // this is openning bracket, push to stack
			stack = append(stack, rString[i])
		}
	}
	return len(stack) == 0
}
