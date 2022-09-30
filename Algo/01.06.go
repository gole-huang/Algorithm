package algo

func SetZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := m, n
	for r := range matrix {
		has0 := false
		for c := range matrix[r] {
			if matrix[r][c] == 0 {
				has0 = true
				break
			}
		}
		if !has0 {
			row = r
			break
		}
	}
	if row == m {
		for r := range matrix {
			for c := range matrix[r] {
				matrix[r][c] = 0
			}
		}
		return
	}
	for c := 0; c < n; c++ {
		has0 := false
		for r := 0; r < m; r++ {
			if matrix[r][c] == 0 {
				has0 = true
				break
			}
		}
		if !has0 {
			col = c
			break
		}
	}
	if col == n {
		for r := range matrix {
			for c := range matrix[r] {
				matrix[r][c] = 0
			}
		}
		return
	}
	for r := range matrix {
		if r == row {
			continue
		}
		for c := range matrix[r] {
			if c != col && matrix[r][c] == 0 {
				matrix[row][c] = 0
				matrix[r][col] = 0
			}
		}
	}
	for r := range matrix {
		if matrix[r][col] == 0 {
			for c := range matrix[r] {
				matrix[r][c] = 0
			}
		}
	}
	for c := range matrix[row] {
		if matrix[row][c] == 0 {
			for r := range matrix {
				matrix[r][c] = 0
			}
		}
	}
}
