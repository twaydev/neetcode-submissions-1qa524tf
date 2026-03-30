func characterReplacement(s string, k int) int {
maxSubstringLength, maxFreq, l := 0, 0, 0
	freq := make(map[byte]int)
	for r := 0; r < len(s); r++ {

		freq[s[r]]++
		maxFreq = max(maxFreq, freq[s[r]])
		if r-l+1-maxFreq > k {
			freq[s[l]]--
			l++
		}
		maxSubstringLength = max(maxSubstringLength, r-l+1)
	}

	return maxSubstringLength
}
