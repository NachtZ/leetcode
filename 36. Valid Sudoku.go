
func isValidSudoku(board [][]byte) bool {
	var rows, cols, blocks [9][9]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			rows[i][j] = false
			cols[i][j] = false
			blocks[i][j] = false
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - 49
			if rows[i][num] || cols[j][num] || blocks[i-i%3+j/3][num] {
				return false
			}
			rows[i][num] = true
			cols[j][num] = true
			blocks[i-i%3+j/3][num] = true
		}
	}
	return true

}