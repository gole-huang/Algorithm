package algo

func pathSum(x, y int, grid [][]int) int {
	if x == len(grid)-1 && y == len(grid[x])-1 {
		return grid[x][y]
	} else if x == len(grid)-1 {
		res := grid[x][y]
		for i := 1; i < len(grid[x])-y; i++ {
			res += grid[x][y+i]
		}
		return res
	} else if y == len(grid[x])-1 {
		res := grid[x][y]
		for i := 1; i < len(grid)-x; i++ {
			res += grid[x+i][y]
		}
		return res
	} else {
		res1, res2 := pathSum(x+1, y, grid), pathSum(x, y+1, grid)
		if res1 > res2 {
			return res2
		} else {
			return res1
		}
	}
}

func MinPathSum(grid [][]int) int {
	return pathSum(0, 0, grid)
}
