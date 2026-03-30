func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	l, r := 0, rows*cols-1

	for l <= r {
		m := l + (r-l)/2
		row := m / cols
		col := m % cols

		if target == matrix[row][col] {
			return true
		} else if target < matrix[row][col] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}
