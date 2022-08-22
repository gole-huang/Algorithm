package Algo

func addLine(sqrs [][3]int, line []byte, max *int) [][3]int {
	newSqrs := make([][3]int, 0) // 存放当前行新增的矩阵（高度为1）
	newSqr := [3]int{1, -1, -1}  // 存放每一个单元格可能新增的矩阵（高度为1）
	for keyLine, vLine := range line {
		tmpSqr := make([][3]int, 0) // 存放之前仍未终结的矩阵
		if vLine == '0' {
			for kCur := 0; kCur < len(sqrs); kCur++ { // 若节点为0，则终结越过该节点的矩阵
				if sqrs[kCur][1] <= keyLine && keyLine < sqrs[kCur][2] {
					if *max < sqrs[kCur][0]*(sqrs[kCur][2]-sqrs[kCur][1]) {
						*max = sqrs[kCur][0] * (sqrs[kCur][2] - sqrs[kCur][1]) // 求当前矩阵最大值
					}
					if keyLine < sqrs[kCur][2]-1 {
						tmpSqr = append(tmpSqr, [3]int{sqrs[kCur][0], keyLine + 1, sqrs[kCur][2]})
					}
					if sqrs[kCur][1] == keyLine {
						sqrs = append(sqrs[:kCur], sqrs[kCur+1:]...)
						kCur--
					} else {
						sqrs[kCur][2] = keyLine
						if sqrs[kCur][1] == newSqr[1] { // 新线段起始与结束与新线段同样，删除新线段.
							newSqr = [3]int{1, -1, -1}
						}
					}
				}
			}
			if newSqr[1] != -1 && newSqr[2] == -1 { //	更新当前行新增矩阵（实际是一条线段）的结束位置
				newSqr[2] = keyLine
				newSqrs = append(newSqrs, newSqr)
				newSqr = [3]int{1, -1, -1}
			}
			sqrs = append(sqrs, tmpSqr...)
		} else if newSqr[1] == -1 {
			newSqr[1] = keyLine
		}
	}

	for k := range sqrs {
		sqrs[k][0]++ // 所有现存的矩形高度+1
	}
	if newSqr[1] != -1 && newSqr[2] == -1 { //	更新当前行新增矩形（实际是一条线段）的结束位置
		newSqr[2] = len(line)
		newSqrs = append(newSqrs, newSqr)
	}
	sqrs = append(sqrs, newSqrs...)
	return sqrs
}

func MaximalRectangle(matrix [][]byte) int {
	maxSquare := new(int)
	curSquare := make([][3]int, 0)
	left := -1
	for k, v := range matrix[0] {
		if v != '0' && left == -1 {
			left = k
		}
		if v == '0' && left != -1 {
			curSquare = append(curSquare, [3]int{1, left, k}) // 左开放，右关闭
			left = -1
		}
	}
	if left != -1 {
		curSquare = append(curSquare, [3]int{1, left, len(matrix[0])}) // 左开放，右关闭
	}
	for _, v := range matrix[1:] {
		curSquare = addLine(curSquare, v, maxSquare)
	}
	for _, v := range curSquare {
		if *maxSquare < v[0]*(v[2]-v[1]) {
			*maxSquare = v[0] * (v[2] - v[1])
		}
	}
	return *maxSquare
}
