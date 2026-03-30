func isValid(s string) bool {
    
	tmp := map[rune]rune{')': '(', '}': '{', ']': '['}
	stack := []rune{} // FILO queue
	for _, c := range s {
		// this is closing bracket, look up for opening bracket
		if _, exist := tmp[c]; exist {
			if len(stack) > 0 && stack[len(stack)-1] == tmp[c] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else { // this is openning bracket, push to stack
			stack = append(stack, c)
		}
	}
	// if all brackets are found, stack should be empty
	return len(stack) == 0
}
