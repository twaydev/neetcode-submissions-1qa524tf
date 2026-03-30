func groupAnagrams(strs []string) [][]string {
	// tạo hashmap có key là 1 mảng int 26 phần tử và giá trị là 1 mảng string
	// create a hashmap with key is an array with 26 elements and value is a string slice
	vals := make(map[[26]int][]string)

	// duyệt qua mảng strs
	// iterate strs array
	for _, val := range strs {
		// khởi tạo biến count với mảng int 26 phần tử
		// initialize count variable as an array with 26 elements
		count := [26]int{}

		// duyệt qua các kí tự của chuỗi str trong mảng strs
		// iterate through characters of str from strs array
		for i := range val {

			// tăng giá trị của kí tự tìm được trong count
			// increase value of any found character
			count[val[i]-'a']++
		}
		// thêm chuỗi tìm được vào vals với key là count (magic)
		// append val to vals hashmap with count as key
		vals[count] = append(vals[count], val)
	}

	// initialize return slice
	returnVal := [][]string{}

	// iterate vals
	for _, val := range vals {
		// append val to returnVal slice
		returnVal = append(returnVal, val)
	}
	return returnVal
}