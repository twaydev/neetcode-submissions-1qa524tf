func searchMatrix(matrix [][]int, target int) bool {

	for _, row := range matrix {
		// lookup target in row
		if row[0] <= target && target <= row[len(row)-1] {
			// use lower bound
			l, r := 0, len(row)
			for l <= r {
				m := l + (r-l)/2
				if target <= row[m] { // move to left, we must decrease r
					r = m - 1
				} else {
					l = m + 1
				}
			}
			// found target
			if l < len(row) && row[l] == target {
				return true
			}
		}
	}
	return false
}
