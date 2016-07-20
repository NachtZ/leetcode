

func help(board [][]byte, rows, cols, blocks [][]bool, idx int) bool{
	if idx == 81 {
		return true
	}
	i := int(idx / 9)
	j := idx % 9
	if board[i][j] != '.' {
		if help(board, rows, cols, blocks, idx+1){
		    return true
		}
	} else {
		for v := 0; v < 9; v++ {
			if rows[i][v] || cols[j][v] || blocks[i-i%3+j/3][v] {
				continue
			}
			board[i][j] = byte(v + 49)
			rows[i][v] = true
			cols[j][v] = true
			blocks[i-i%3+j/3][v] = true
			if true == help(board, rows, cols, blocks, idx+1){
			    return true
			}
			board[i][j] = '.'
			rows[i][v] = false
			cols[j][v] = false
			blocks[i-i%3+j/3][v] = false
		}
	}
	return false
}

func solveSudoku(board [][]byte) {
	var rows, cols, blocks [][]bool
	for i := 0; i < 9; i++ {
		_rows := []bool{false, false, false, false, false, false, false, false, false}
		_cols := []bool{false, false, false, false, false, false, false, false, false}
		_blocks := []bool{false, false, false, false, false, false, false, false, false}
		rows = append(rows, _rows)
		cols = append(cols, _cols)
		blocks = append(blocks, _blocks)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - 49
			rows[i][num] = true
			cols[j][num] = true
			blocks[i-i%3+j/3][num] = true
		}
	}

	help(board, rows, cols, blocks, 0)

}