func removeNonAlphaAndLowerCase(s string) string {
	var res string
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			res += string(unicode.ToLower(r))
		}
	}
	return res
}

func isPalindrome(s string) bool {

	// remove non alpha
	s = removeNonAlphaAndLowerCase(s)

	// create slice with string
	originSlice := []byte(s)

	// check if palindrome
	// declare i to track from the beginning and j to track from the end
	// check value of slice with i, j
	for i, j := 0, len(originSlice)-1; i < j; i, j = i+1, j-1 {
		if originSlice[i] != originSlice[j] {
			return false
		}
	}
	return true
}