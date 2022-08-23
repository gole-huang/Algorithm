package Algo

func isChesboard(brd [][]int, rowMax, rowMin, colMax, colMin *int) bool {
	one := 0
	for _, b := range brd {
		for _, v := range b {
			if v == 0 {
				one--
				if one < *rowMin {
					*rowMin = one
				}
			} else {
				one++
				if one > *rowMax {
					*rowMax = one
				}
			}
		}
		if one < -1 || one > 1 {
			return false
		}
		one = 0
	}
	for j := 0; j < len(brd[0]); j++ {
		for i := 0; i < len(brd); i++ {
			if brd[i][j] == 0 {
				one--
				if one < *colMin {
					*colMin = one
				}
			} else {
				one++
				if one > *colMax {
					*colMax = one
				}
			}
		}
		if one < -1 || one > 1 {
			return false
		}
		one = 0
	}
	return true
}

func MovesToChessboard(board [][]int) int {
	rowMin, rowMax, colMin, colMax := new(int), new(int), new(int), new(int)
	if !isChesboard(board, rowMin, rowMax, colMin, colMax) {
		return -1
	}
	var row, col int
	if *rowMin+*rowMax < 0 {
		row = -*rowMin
	} else {
		row = *rowMax
	}
	if *colMin+*colMax < 0 {
		col = -*colMin
	} else {
		col = *colMax
	}
	return row/2 + col/2
}
