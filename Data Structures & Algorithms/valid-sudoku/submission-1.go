func isValidSudoku(board [][]byte) bool {
n := len(board)
	// initialize hashsets
	colHashes := make([]map[byte]bool, 9)
	rowHashes := make([]map[byte]bool, 9)
	subSquareHashes := make([]map[byte]bool, 9)
	for i := 0; i < n; i++ {
		colHashes[i] = make(map[byte]bool)
		rowHashes[i] = make(map[byte]bool)
		subSquareHashes[i] = make(map[byte]bool)
	}
	// iterate board to check duplicate values
	for row := 0; row < n; row++ {
		rowHash := rowHashes[row]
		for col := 0; col < n; col++ {
			colHash := colHashes[col]
			if board[row][col] == '.' {
				continue
			}
			// check duplicate on rows
			if _, ok := rowHash[board[row][col]]; ok {
				return false
			}
			// check duplicate on cols
			if _, ok := colHash[board[row][col]]; ok {
				return false
			}

			// calculate sub square index and check duplicate in sub squares
			subSquareIndex := (row/3)*3 + col/3
			subSquareHash := subSquareHashes[subSquareIndex]
			if _, ok := subSquareHash[board[row][col]]; ok {
				return false
			}
			rowHashes[row][board[row][col]] = true
			colHashes[col][board[row][col]] = true
			subSquareHashes[subSquareIndex][board[row][col]] = true
		}
	}
	return true
}
