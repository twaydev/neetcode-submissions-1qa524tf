func hasDuplicate(nums []int) bool {
    
	// nếu mảng có dưới 2 phần tử thì trả về false
	if len(nums) < 2 {
		return false
	}

	// tạo map với dạng key => value kiểm tra phần tu đã thấy
	hasSeen := make(map[int]bool)
	for _, num := range nums {
		if hasSeen[num] == true {
			return true
		} else {
			hasSeen[num] = true
		}
	}

	// trong trường hợp duyệt qua hết mảng mà ko bị trùng lặp
	return false
}
