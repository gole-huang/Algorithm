package Algo

/*
第一行作为比较，之后所有行，要么与第一行一样，要么与第一行相反。列也一样。
第一行中，0的个数和1的个数相差不超过1.同样，第一列中0的个数和1的个数不超过1
若第一行为偶数，则只需要计算出出现连续1的长度和次数。长度/2代表移动1列，把所有长度加起来即为待调整的次数。列也要单独算
若第一行为奇数，则需要判断0和1较多者是否在两个角落，否则每个角落需要额外+1
*/
func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	reverse := a[0] != b[0]
	for k := range a {
		if reverse && a[k] == b[k] || !reverse && a[k] != b[k] {
			return false
		}
	}
	return true
}

func isChessboard(brd [][]int) (bool, int) {
	zero, one := 0, 0   // 首行、首列的0、1个数
	con0, con1 := 1, 1  // 每一个连续0，1的长度
	zeros, ones := 0, 0 // 首行、首列连续0、1中需要移动的次数
	var cnt int         //交换总数
	//首行
	for i := 0; i < len(brd); i++ {
		if brd[0][i] == 0 {
			zero++
			if i > 0 {
				if brd[0][i-1] == 0 {
					con0++
				} else if con1 > 1 {
					ones += con1 >> 1
					con1 = 1
				}
			}
		} else {
			one++
			if i > 0 {
				if brd[0][i-1] == 1 {
					con1++
				} else {
					zeros += con0 >> 1
					con0 = 1
				}
			}
		}

	}
	if len(brd)%2 == 0 {
		if zero != one {
			return false, -1
		}
		zeros += con0 >> 1
		ones += con1 >> 1
		cnt = max(zeros, ones)
	} else if zero > one {
		if zero-one != 1 {
			return false, -1
		}
		if brd[0][0] != 0 {
			zeros++
		}
		if brd[len(brd)-1][0] != 0 {
			zeros++
		}
		cnt = zeros
	} else {
		if one-zero != 1 {
			return false, -1
		}
		if brd[0][0] != 1 {
			ones++
		}
		if brd[len(brd)-1][0] != 0 {
			ones++
		}
		cnt = ones
	}
	zeros, one = 0, 0
	for i := 0; i < len(brd); i++ {
		if brd[i][0] == 0 {
			zero++
			if i > 0 {
				if brd[0][i-1] == 0 {
					con0++
				} else if con1 > 1 {
					ones += con1 >> 1
					con1 = 1
				}
			}
		} else {
			one++
			if i > 0 {
				if brd[0][i-1] == 1 {
					con1++
				} else {
					zeros += con0 >> 1
					con0 = 1
				}
			}
		}
	}
	if len(brd)%2 == 0 {
		if zero != one {
			return false, -1
		}
		zeros += con0 >> 1
		ones += con1 >> 1
		cnt += max(zeros, ones)
	} else if zero > one {
		if zero-one != 1 {
			return false, -1
		}
		if brd[0][0] != 0 {
			zeros++
		}
		if brd[len(brd)-1][0] != 0 {
			zeros++
		}
		cnt += zeros
	} else {
		if one-zero != 1 {
			return false, -1
		}
		if brd[0][0] != 1 {
			ones++
		}
		if brd[len(brd)-1][0] != 0 {
			ones++
		}
		cnt += ones
	}
	return true, max(zeros, ones)
}

func movesToChessboard(board [][]int) int {
	firstCheck, res := isChessboard(board)
	if !firstCheck {
		return -1
	}
	for i := 1; i < len(board); i++ {
		if !compare(board[0], board[i]) {
			return -1
		}
	}
}
