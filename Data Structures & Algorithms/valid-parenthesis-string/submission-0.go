

func checkValidString(s string) bool {
	parenthesisStack := []int{} // Store INDICES, not characters
	asteriskStack := []int{}    // Store INDICES

	for i, c := range s {
		switch c {
		case '(':
			parenthesisStack = append(parenthesisStack, i)
		case ')':
			if len(parenthesisStack) > 0 {
				parenthesisStack = parenthesisStack[:len(parenthesisStack)-1]
			} else if len(asteriskStack) > 0 {
				asteriskStack = asteriskStack[:len(asteriskStack)-1]
			} else {
				return false
			}
		default: // '*'
			asteriskStack = append(asteriskStack, i)
		}
	}

	// Match remaining '(' with '*' that come AFTER them
	for len(parenthesisStack) > 0 && len(asteriskStack) > 0 {
		if parenthesisStack[len(parenthesisStack)-1] > asteriskStack[len(asteriskStack)-1] {
			// '(' comes after '*', can't use '*' as ')'
			return false
		}
		parenthesisStack = parenthesisStack[:len(parenthesisStack)-1]
		asteriskStack = asteriskStack[:len(asteriskStack)-1]
	}

	return len(parenthesisStack) == 0
}
