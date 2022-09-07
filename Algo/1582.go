package Algo

func numSpecial(mat [][]int) int {
	res := 0
	for kr := range mat {
		rowOnce, rowTwice := false, false
		colTwice := false
		for kc := range mat[kr] {
			if mat[kr][kc] == 1 {
				if !rowOnce {
					rowOnce = true
				} else {
					rowTwice = true
					break
				}
				for k := range mat {
					if k == kr {
						continue
					} else if mat[k][kc] == 1 {
						colTwice = true
						break
					}
				}
			}
		}
		if rowOnce && !rowTwice && !colTwice {
			res++
		}
	}
	return res
}
