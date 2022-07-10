package Algo

import "fmt"

const DMS int = 9

type Sudoku struct {
	brd  [][]byte
	rows []map[byte]bool
	cols []map[byte]bool
	mtxs []map[byte]bool
}

/*TODO
为数独的每一行、每一列、每一个矩阵计算出可用数字，分别放入rows/cols/mtxs中；
当前空格尝试填入一个数字，该数字必须为该行、列、矩阵中的可用数字，且非该行、列、矩阵的已占用数字。填入后该数字添加到该行、列、矩阵的已占用数字中；
不断尝试填入并寻找下一个空位，一直到填完整个数独为止（下一个空位坐标为（-1，-1）），返回真，若没有办法完成，返回假
*/

func (sdk *Sudoku) rowValid() {
	sdk.rows = make([]map[byte]bool, DMS)
	for row := 0; row < DMS; row++ {
		sdk.rows[row] = make(map[byte]bool)
		for val := 0; val < DMS; val++ {
			sdk.rows[row][byte(49+val)] = true
		}
	}
	for row := 0; row < DMS; row++ {
		for col := 0; col < DMS; col++ {
			if sdk.brd[row][col] != '.' {
				delete(sdk.rows[row], sdk.brd[row][col])
			}
		}
	}
}

func (sdk *Sudoku) colValid() {
	sdk.cols = make([]map[byte]bool, DMS)
	for col := 0; col < DMS; col++ {
		sdk.cols[col] = make(map[byte]bool)
		for val := 0; val < DMS; val++ {
			sdk.cols[col][byte(49+val)] = true
		}
	}
	for col := 0; col < DMS; col++ {
		for row := 0; row < DMS; row++ {
			if sdk.brd[row][col] != '.' {
				delete(sdk.cols[col], sdk.brd[row][col])
			}
		}
	}
}

func (sdk *Sudoku) sqrValid() {
	sdk.mtxs = make([]map[byte]bool, DMS)
	for mtx := 0; mtx < DMS; mtx++ {
		sdk.mtxs[mtx] = make(map[byte]bool)
		for val := 0; val < DMS; val++ {
			sdk.mtxs[mtx][byte(49+val)] = true
		}
	}
	for row := 0; row < DMS; row++ {
		for col := 0; col < DMS; col++ {
			mtx := (row/3)*3 + (col / 3)
			if sdk.brd[row][col] != '.' {
				delete(sdk.mtxs[mtx], sdk.brd[row][col])
			}
		}
	}
}

func (sdk *Sudoku) findNext(x, y int) (int, int) {
	for j := y; j < DMS; j++ {
		if sdk.brd[x][j] == '.' {
			return x, j
		}
	}
	for i := x + 1; i < DMS; i++ {
		for j := 0; j < DMS; j++ {
			if sdk.brd[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func (sdk *Sudoku) fill(x, y int, rowUsed, colUsed, mtxUsed []map[byte]bool) bool {
	if x == -1 && y == -1 {
		return true
	}
	sqr := (x/3)*3 + (y / 3)
	for i := uint8(49); i < uint8(58); i++ {
		if sdk.rows[x][i] && !rowUsed[x][i] && sdk.cols[y][i] && !colUsed[y][i] && sdk.mtxs[sqr][i] && !mtxUsed[sqr][i] {
			rowUsed[x][i] = true
			colUsed[y][i] = true
			mtxUsed[sqr][i] = true
			fmt.Printf("%d,%d, try i = %q\n", x, y, i)
			xNext, yNext := sdk.findNext(x, y+1)
			if !sdk.fill(xNext, yNext, rowUsed, colUsed, mtxUsed) {
				rowUsed[x][i] = false
				colUsed[y][i] = false
				mtxUsed[sqr][i] = false
				continue
			} else {
				sdk.brd[x][y] = i
				return true
			}
		}
	}
	return false
}

func SolveSudoku(board [][]byte) {
	sudoku := new(Sudoku)
	sudoku.brd = board
	//初始化可用数字
	sudoku.rowValid()
	sudoku.colValid()
	sudoku.sqrValid()
	//初始化被占用数字
	rowUsed := make([]map[byte]bool, DMS)
	colUsed := make([]map[byte]bool, DMS)
	mtxUsed := make([]map[byte]bool, DMS)
	for i := 0; i < DMS; i++ {
		rowUsed[i] = make(map[byte]bool)
		colUsed[i] = make(map[byte]bool)
		mtxUsed[i] = make(map[byte]bool)
	}
	x, y := sudoku.findNext(0, 0)
	if sudoku.fill(x, y, rowUsed, colUsed, mtxUsed) {
		board = sudoku.brd
	}
	fmt.Printf("%v\n", board)
}
