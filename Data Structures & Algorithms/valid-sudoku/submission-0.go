func isValidSudoku(board [][]byte) bool {

	n := len(board)

	// hashset must belong to a row, a col and a square
	rowsHash := make([]map[byte]bool, 9) // len=9 and cap=9
	for r := range rowsHash {
		rowsHash[r] = make(map[byte]bool)
	}

	// check row
	for i := 0; i < n; i++ {
		rowHash := rowsHash[i] //make(map[byte]bool)
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := rowHash[board[i][j]]; ok {
				return false
			}
			rowsHash[i][board[i][j]] = true
		}
	}

	// check col
	colsHash := make([]map[byte]bool, 9)
	for c := range colsHash {
		colsHash[c] = make(map[byte]bool)
	}
	for i := 0; i < n; i++ {
		colHash := colsHash[i]
		for j := 0; j < n; j++ {
			if board[j][i] == '.' {
				continue
			}
			if _, ok := colHash[board[j][i]]; ok {
				return false
			}
			colsHash[i][board[j][i]] = true
		}
	}

	// check sub square
	squaresHash := make([]map[byte]bool, 9)
	for s := range squaresHash {
		squaresHash[s] = make(map[byte]bool)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}
			squareIndex := (i/3)*3 + (j / 3)
			if _, ok := squaresHash[squareIndex][board[i][j]]; ok {
				return false
			}
			squaresHash[squareIndex][board[i][j]] = true
		}
	}

	return true
}
