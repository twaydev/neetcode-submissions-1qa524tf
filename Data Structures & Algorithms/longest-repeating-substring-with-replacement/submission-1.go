func characterReplacement(s string, k int) int {

	maxSubstringLength, maxFreq, l := 0, 0, 0
	freq := make(map[byte]int)
	for r := 0; r < len(s); r++ {

		freq[s[r]]++
		maxFreq = max(maxFreq, freq[s[r]])
		windowSize := r - l + 1
		if windowSize-maxFreq > k {
			freq[s[l]]--
			l++
			windowSize--
		}
		maxSubstringLength = max(maxSubstringLength, windowSize)
	}

	return maxSubstringLength
}
