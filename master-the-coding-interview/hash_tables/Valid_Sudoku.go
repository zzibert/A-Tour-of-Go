package main

func isValidSudoku(board [][]byte) bool {

	// check if rows are valid.
	for i := 0; i < 9; i++ {
		if !isValidRow(board[i]) {
			return false
		}
	}

	// check if columns are valid.
	for i := 0; i < 9; i++ {
		if !isValidColumn(board, i) {
			return false
		}
	}

	// check if subgrids are valid.
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !isValidSubGrid(board, i, j) {
				return false
			}
		}
	}

	return true

}

func isValidRow(row []byte) bool {
	numbers := make(map[byte]bool)

	for _, number := range row {
		if number != '.' && numbers[number] {
			return false
		}
		numbers[number] = true
	}

	return true
}

func isValidColumn(board [][]byte, column int) bool {
	numbers := make(map[byte]bool)

	for i := 0; i < 9; i++ {
		if board[i][column] != '.' && numbers[board[i][column]] {
			return false
		}
		numbers[board[i][column]] = true
	}

	return true
}

func isValidSubGrid(board [][]byte, row, column int) bool {
	numbers := make(map[byte]bool)

	for i := row; i < row+3; i++ {
		for j := column; j < column+3; j++ {
			if board[i][j] != '.' && numbers[board[i][j]] {
				return false
			}
			numbers[board[i][j]] = true
		}
	}

	return true
}
