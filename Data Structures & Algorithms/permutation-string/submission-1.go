
func checkEqual(m1, m2 map[byte]int) bool {
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	n := len(s1)
	// delcare char map to tracking freq of current window
	// delcare char map to tracking freq of s1
	windowCharMap, s1CharMap := make(map[byte]int, 26), make(map[byte]int, 26)
	for i := range n {
		windowCharMap[s2[i]-'a']++
		s1CharMap[s1[i]-'a']++
	}

	// delcare left and right
	l, r := 0, n-1
	for r < len(s2) {
		// fmt.Println(l, r, s1CharMap, windowCharMap, string(s2[l]), isEqual)
		if isEqual := checkEqual(s1CharMap, windowCharMap); isEqual {
			return true
		}
		// not equal,  decrease freq of left and increase left
		windowCharMap[s2[l]-'a']--
		l++
		// increase right
		r++
		if r > len(s2)-1 {
			break
		}
		windowCharMap[s2[r]-'a']++
	}
	return false
}