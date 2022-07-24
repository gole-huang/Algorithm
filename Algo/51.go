package Algo

func setQueen(grid [][]bool, x, y int, oCols []int) (res [][][2]int) {
	n := len(grid)
	if grid[x][oCols[y]] {
		return nil
	} else if x == n-1 {
		return [][][2]int{{{x, oCols[y]}}}
	}
	tmpGrid := make([][]bool, n)
	for i := 0; i < n; i++ {
		tmpGrid[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			ix, jy := i-x, j-oCols[y]
			if i == x || j == oCols[y] || ix == jy || ix+jy == 0 {
				tmpGrid[i][j] = true
			} else {
				tmpGrid[i][j] = grid[i][j]
			}
		}
	}
	cols := make([]int, len(oCols))
	copy(cols, oCols)
	cols = append(cols[:y], cols[y+1:]...)
	for i := range cols {
		if !tmpGrid[x+1][cols[i]] {
			tmpRes := setQueen(tmpGrid, x+1, i, cols)
			for tmp := range tmpRes {
				tmpRes[tmp] = append([][2]int{{x, oCols[y]}}, tmpRes[tmp]...)
			}
			res = append(res, tmpRes...)
		}
	}
	return res
}

func SolveNQueens(n int) [][]string {
	cols := make([]int, n)
	grid := make([][]bool, n)
	for i := range grid {
		cols[i] = i
		grid[i] = make([]bool, n)
	}
	intRes := make([][][2]int, 0)
	for i := 0; i < n; i++ {
		tmpRes := setQueen(grid, 0, i, cols)
		if tmpRes != nil {
			intRes = append(intRes, tmpRes...)
		}
		// fmt.Printf("intRes=%v\n", intRes)
	}
	strRes := make([][]string, 0)
	for i := 0; i < len(intRes); i++ {
		str := make([]string, 0)
		for j := 0; j < n; j++ {
			byteRes := make([]byte, n)
			for k := 0; k < n; k++ {
				if j == intRes[i][j][0] && k == intRes[i][j][1] {
					byteRes[k] = 'Q'
				} else {
					byteRes[k] = '.'
				}
			}
			str = append(str, string(byteRes))
		}
		strRes = append(strRes, str)
	}
	return strRes
}