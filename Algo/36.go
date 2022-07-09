package Algo

type Sudoku struct {
	brd  [][]byte
	rows []map[byte]bool
	cols []map[byte]bool
	sqrs []map[byte]bool
}

func (sdk *Sudoku) rowValid() {
	sdk.rows = make([]map[byte]bool, len(sdk.brd))
	for row := 0; row < len(sdk.brd); row++ {
		for val := 0; val < len(sdk.brd[row]); val++ {
			sdk.rows[row][byte(val+1)] = true
		}
	}
	for row := 0; row < len(sdk.brd); row++ {
		for col := 0; col < len(sdk.brd[row]); col++ {
			if sdk.brd[row][col] != '.' {
				delete(sdk.rows[row], sdk.brd[row][col])
			}
		}
	}
}

func (sdk *Sudoku) colValid() {
	sdk.cols = make([]map[byte]bool, len(sdk.brd[0]))
	for col := 0; col < len(sdk.brd); col++ {
		for val := 0; val < len(sdk.brd[col]); val++ {
			sdk.cols[col][byte(val+1)] = true
		}
	}
	for col := 0; col < len(sdk.brd[0]); col++ {
		for row := 0; row < len(sdk.brd); row++ {
			if sdk.brd[row][col] != '.' {
				delete(sdk.cols[col], sdk.brd[row][col])
			}
		}
	}
}

func (sdk *Sudoku) sqrValid() {
	sdk.sqrs = make([]map[byte]bool, len(sdk.brd[0]))
	for sqr := 0; sqr < len(sdk.brd); sqr++ {
		for val := 0; val < len(sdk.brd[sqr]); val++ {
			sdk.sqrs[sqr][byte(val+1)] = true
		}
	}
	for row := 0; row < len(sdk.brd); row++ {
		for col := 0; col < len(sdk.brd[0]); col++ {
			sqr := row/3 + (col/3)*3
			if sdk.brd[row][col] != '.' {
				delete(sdk.sqrs[sqr], sdk.brd[row][col])
			}
		}
	}
}

func (sdk *Sudoku) findNext(x, y int) (int, int) {
	for j := y; j < len(sdk.brd[x]); j++ {
		if sdk.brd[x][j] == '.' {
			return x, j
		}
	}
	for i := x + 1; i < len(sdk.brd); i++ {
		for j := 0; j < len(sdk.brd[i]); j++ {
			if sdk.brd[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func (sdk *Sudoku) fill(x, y int, rowUsed, colUsed, sqrUsed []map[byte]bool) bool {
	if x == -1 && y == -1 {
		return true
	}
	sqr := x/3 + (y/3)*3
	for i := uint8(49); i < uint8(59); i++ {
		if sdk.rows[x][i] && !rowUsed[x][i] && sdk.cols[y][i] && !colUsed[y][i] && sdk.sqrs[sqr][i] && !sqrUsed[sqr][i] {
			sdk.brd[x][y] = i
			rowUsed[x][i] = true
			colUsed[y][i] = true
			sqrUsed[sqr][i] = true
			xNext, yNext := sdk.findNext(x, y+1)
			return sdk.fill(xNext, yNext, rowUsed, colUsed, sqrUsed)
		}
	}
	return false
}

func IsValidSudoku(board [][]byte) bool {
	sudoku := new(Sudoku)
	sudoku.brd = board
	sudoku.rowValid()
	sudoku.colValid()
	sudoku.sqrValid()
	rowUsed := make([]map[byte]bool, len(sudoku.brd))
	colUsed := make([]map[byte]bool, len(sudoku.brd[0]))
	sqrUsed := make([]map[byte]bool, len(sudoku.brd)/3+(len(sudoku.brd[0])/3)*3)
	return sudoku.fill(0, 0, rowUsed, colUsed, sqrUsed)
}
