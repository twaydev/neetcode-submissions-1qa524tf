func isNStraightHand(hand []int, groupSize int) bool {
    // sort hand asc
	sort.Ints(hand)

	// count freq
	freq := make(map[int]int)
	for _, num := range hand {
		freq[num]++
	}

	// Duyệt và tạo nhóm
	for _, card := range hand {
		if freq[card] == 0 {
			continue // Đã dùng hết
		}

		// Cố gắng tạo nhóm từ card
		for i := range groupSize {
			current := card + i
			if freq[current] == 0 {
				return false
			}
			freq[current]--
		}
	}

	return true
}
