package Algo

func addLine(squares map[[2]int]int, line []byte, max *int) map[[2]int]int {
	newLines := make([][2]int, 0) // 存放当前行新增的矩阵（高度为1）
	newLine := [2]int{-1, -1}     // 存放每一个单元格可能新增的矩阵（高度为1）
	for keyLine, vLine := range line {
		if vLine == '0' {
			for key, value := range squares { // 若节点为0，则终结越过该节点的矩阵
				if key[0] <= keyLine && keyLine < key[1] {
					if *max < value*(key[1]-key[0]) {
						*max = value * (key[1] - key[0]) // 求当前矩阵最大值
					}
					if keyLine < key[1]-1 {
						if v, ok := squares[[2]int{keyLine + 1, key[1]}]; !ok || v < value {
							squares[[2]int{keyLine + 1, key[1]}] = value
						}
						delete(squares, key)
					}
					if key[0] == keyLine {
						delete(squares, key)
					} else {
						if v, ok := squares[[2]int{key[0], keyLine}]; !ok || v < value {
							squares[[2]int{key[0], keyLine}] = value
						}
						if key[0] == newLine[0] { // 新线段起始与结束与新线段同样，删除新线段.
							newLine = [2]int{-1, -1}
						}
						delete(squares, key)
					}
				}
			}
			if newLine[0] != -1 && newLine[1] == -1 { //	更新当前行新增矩阵（实际是一条线段）的结束位置
				newLine[1] = keyLine
				newLines = append(newLines, newLine)
				newLine = [2]int{-1, -1}
			}
		} else if newLine[0] == -1 {
			newLine[0] = keyLine
		}
	}

	for k := range squares {
		squares[k]++ // 所有现存的矩形高度+1
	}
	if newLine[0] != -1 && newLine[1] == -1 { //	更新当前行新增矩形（实际是一条线段）的结束位置
		newLine[1] = len(line)
		newLines = append(newLines, newLine)
	}
	for _, v := range newLines {
		if _, ok := squares[[2]int{v[0], v[1]}]; !ok {
			squares[[2]int{v[0], v[1]}] = 1
		}
	}
	return squares
}

func MaximalRectangle(matrix [][]byte) int {
	maxSquare := new(int)
	curSquare := make(map[[2]int]int)
	left := -1
	for k, v := range matrix[0] {
		if v != '0' && left == -1 {
			left = k
		}
		if v == '0' && left != -1 {
			curSquare[[2]int{left, k}] = 1 // 左开放，右关闭
			left = -1
		}
	}
	if left != -1 {
		curSquare[[2]int{left, len(matrix[0])}] = 1 // 左开放，右关闭
	}
	for _, line := range matrix[1:] {
		curSquare = addLine(curSquare, line, maxSquare)
	}
	for k, v := range curSquare {
		if *maxSquare < v*(k[1]-k[0]) {
			*maxSquare = v * (k[1] - k[0])
		}
	}
	return *maxSquare
}
