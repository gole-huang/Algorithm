package algo

// 第一行作为比较，之后所有行，要么与第一行一样，要么与第一行相反。
// 列也一样。
func compare(a, b []int) bool {
	reverse := a[0] != b[0]
	for k := range a {
		if reverse && a[k] == b[k] || !reverse && a[k] != b[k] {
			return false
		}
	}
	return true
}

func isChessboard(brd [][]int) (bool, int) {
	_0to0, _0to1, _1to0, _1to1 := 0, 0, 0, 0
	var cnt int // 交换总数

	for i := range brd[0] { // 首行
		if brd[0][i] == 0 {
			if i%2 != 0 { // 假设棋盘首行为0101010101......，不是也无所谓，比较每一位与该位置棋盘该有的数字做比较。
				_0to1++ // 若应该为0（i%2==0），而此处为1，则_0to1++
			} else {
				_0to0++ // 若应该为0（i%2==0），而此处为0，则_0to0++
			}
		} else {
			if i%2 != 1 {
				_1to0++ // 若应该为0（i%2==1），而此处为1，则_1to0++
			} else {
				_1to1++ // 若应该为1（i%2==1），而此处为1，则_1to1++
			}
		}
	}
	if len(brd)%2 == 0 { // 若n为偶数，则0到1的转换应该和1到0的转换一样。如果棋盘是10101010......，则0到0的转换应该等于1到1的转换
		if _0to1 != _1to0 || _0to0 != _1to1 {
			return false, -1
		}
		if _0to1+_1to0 > _0to0+_1to1 {
			cnt += _0to0
		} else {
			cnt += _0to1
		}
	} else { //若n为偶数，则情况比较复杂。要么0到1的转换等于1到0的转换，且0到0的转换和1到1的转换之差不超过1，或者相反
		if (_0to1 != _1to0 || (_0to0-_1to1 != 1 && _0to0-_1to1 != -1)) && (_0to0 != _1to1 || (_0to1-_1to0 != 1 && _0to1-_1to0 != -1)) {
			return false, -1
		}
		if _0to1 != _1to0 {
			cnt += _0to0
		} else {
			cnt += _0to1
		}
	}
	_0to0, _0to1, _1to0, _1to1 = 0, 0, 0, 0
	for i := range brd { // 首列
		if brd[i][0] == 0 {
			if i%2 != 0 {
				_0to1++
			} else {
				_0to0++
			}
		} else {
			if i%2 != 1 {
				_1to0++
			} else {
				_1to1++
			}
		}
	}
	if len(brd)%2 == 0 {
		if _0to1 != _1to0 || _0to0 != _1to1 {
			return false, -1
		}
		if _0to1+_1to0 > _0to0+_1to1 {
			cnt += _0to0
		} else {
			cnt += _0to1
		}
	} else {
		if (_0to1 != _1to0 || (_0to0-_1to1 != 1 && _0to0-_1to1 != -1)) && (_0to0 != _1to1 || (_0to1-_1to0 != 1 && _0to1-_1to0 != -1)) {
			return false, -1
		}
		if _0to1 != _1to0 {
			cnt += _0to0
		} else {
			cnt += _0to1
		}
	}
	return true, cnt
}

func MovesToChessboard(board [][]int) int {
	firstCheck, res := isChessboard(board)
	if !firstCheck {
		return -1
	}
	for i := 1; i < len(board); i++ {
		if !compare(board[0], board[i]) {
			return -1
		}
	}
	for i := 1; i < len(board); i++ {
		reverse := board[i][0] != board[0][0]
		for k := range board[i] {
			if reverse && board[i][k] == board[0][k] || !reverse && board[i][k] != board[0][k] {
				return -1
			}
		}
	}
	return res
}
