
func longestPalindrome(s string) string {
	n := len(s)
	subLen := 0
	subIdx := 0

	expand := func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > subLen {
				subLen = r - l + 1
				subIdx = l
			}
			l--
			r++
		}
	}

	for i := range n {
		expand(i, i)   // odd-length palindrome centered at i
		expand(i, i+1) // even-length palindrome centered between i and i+1
	}

	return s[subIdx : subIdx+subLen]
}
