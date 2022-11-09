package algo

func OrderOfLargestPlusSign(n int, mines [][]int) int {
	sqr := make([][]int, n)
	for k := range sqr {
		sqr[k] = make([]int, n)
	}
	for _, v := range mines {
		sqr[v[0]][v[1]] = 1
	}
	maxCross := 0
	for i := range sqr {
		for j := range sqr[i] {
			edge := func() int {
				var a, b int
				if i < n/2 {
					a = i
				} else {
					a = n - 1 - i
				}
				if j < n/2 {
					b = j
				} else {
					b = n - 1 - j
				}
				if a < b {
					return a
				} else {
					return b
				}
			}
			if edge() < maxCross {
				continue
			}
			curCross := 0
			for k := 0; k <= edge(); k++ {
				if sqr[i-k][j] == 1 || sqr[i+k][j] == 1 || sqr[i][j-k] == 1 || sqr[i][j+k] == 1 {
					break
				}
				curCross++
			}
			if maxCross < curCross {
				maxCross = curCross
			}
		}
	}
	return maxCross
}
