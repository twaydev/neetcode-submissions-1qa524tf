func lengthOfLongestSubstring(s string) int {
if len(s) == 0 {
		return 0
	}
	charIndex := make(map[byte]int)
	left := 0
	maxLength := 0
	for right := 0; right < len(s); right++ {
		if idx, exist := charIndex[s[right]]; exist && idx >= left {
			left = idx + 1
		}
		charIndex[s[right]] = right // save curent char index to check it
		// update maxLength
		max := right - left + 1
		if max > maxLength {
			maxLength = max
		}
	}
	return maxLength
}
