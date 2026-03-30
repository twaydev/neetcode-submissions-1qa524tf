
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	n := len(s1)
	// delcare char map to tracking freq of current window
	// delcare char map to tracking freq of s1
	wCharArray, s1CharArray := [26]int{}, [26]int{}
	for i := range n {
		wCharArray[s2[i]-'a']++
		s1CharArray[s1[i]-'a']++
	}
	equal := func() bool {
		for i := range 26 {
			if wCharArray[i] != s1CharArray[i] {
				return false
			}
		}
		return true
	}
	if equal() {
		return true
	}
	for i := n; i < len(s2); i++ {
		wCharArray[s2[i]-'a']++
		wCharArray[s2[i-n]-'a']--
		if equal() {
			return true
		}
	}
	return false
}