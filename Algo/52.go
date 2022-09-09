package algo

/*
TotalNQueens
The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.Given an integer n, return the number of distinct solutions to the n-queens puzzle.
思路：
res：每次完成解答，自身加1
递归函数：参数为当前当前行，当前可用列、当前可用棋盘。
方法为从可用列取出一个值，与当前行组成一对坐标。
-若该坐标不可用则继续下一列
-若该坐标可用，则更新可用列和可用棋盘，返回递归函数，行为下一行，可用列减去当前列，可用棋盘为更新棋盘
--更新棋盘的方法为，当前行置为不可用，当前列置为不可用，当前行、当前列同时+i、同时-i、分别+-i不可用
-若所有可用列均不可用，则返回false
完成解答：最后一个放置棋盘位置可用，res++
*/
func setQueen2(x uint8, colsValid []uint8, oldBoard [][]bool, res *int) {
	if len(colsValid) == 1 {
		if oldBoard[x][colsValid[0]] {
			*res++
		}
	} else {
		for c := range colsValid {
			if oldBoard[x][colsValid[c]] {
				cols := make([]uint8, len(colsValid))
				copy(cols, colsValid)
				cols = append(cols[:c], cols[c+1:]...)
				board := make([][]bool, len(oldBoard))
				for i := 0; i < len(board); i++ {
					board[i] = make([]bool, len(oldBoard[i]))
					for j := 0; j < len(board[i]); j++ {
						if i == int(x) || j == int(colsValid[c]) || int(x)-i == int(colsValid[c])-j || int(x+colsValid[c]) == i+j {
							board[i][j] = false
						} else {
							board[i][j] = oldBoard[i][j]
						}
					}
				}
				setQueen2(x+1, cols, board, res)
			}
		}
	}
}
func TotalNQueens(n int) int {
	res := new(int)
	*res = 0
	cols := make([]byte, n)
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		cols[i] = byte(i)
		board[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			board[i][j] = true
		}
	}
	setQueen2(0, cols, board, res)
	return *res
}
