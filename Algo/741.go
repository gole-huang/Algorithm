package Algo

type farm struct {
	grid   [][]int
	length int
}

func (cherry *farm) pickCount(x, y int) int {
	if cherry.grid[x][y] == -1 {
		return -1
	} else {
		res := cherry.grid[x][y]
		cherry.grid[x][y] = 0
		if x == cherry.length-1 {
			for i := y; i < cherry.length; i++ {
				if cherry.grid[x][i] == -1 {
					return -1
				} else if cherry.grid[x][i] == 1 {
					res++
					cherry.grid[x][i] = 0
				}
			}
			return res
		} else if y == cherry.length-1 {
			for i := x; i < cherry.length-1; i++ {
				if cherry.grid[i][y] == -1 {
					return -1
				} else if cherry.grid[i][y] == 1 {
					res++
					cherry.grid[i][y] = 0
				}
			}
			return res
		} else {
			xMove := cherry.pickCount(x+1, y)
			yMove := cherry.pickCount(x, y+1)
			if xMove == -1 && yMove == -1 {
				return -1
			} else if xMove > yMove {
				return res + xMove
			} else {
				return res + yMove
			}
		}
	}
}

/*
func (cherry *farm) upCount(x, y int) int {
	if cherry.grid[x][y] == -1 {
		return -1
	} else {
		if x == 0 {
			res := 0
			for i := y; i >= 0; i-- {
				if cherry.grid[x][i] == -1 {
					return -1
				} else {
					res += cherry.grid[0][i]
				}
			}
			return res
		} else if y == 0 {
			res := 0
			for i := x; i >= 0; i-- {
				if cherry.grid[i][y] == -1 {
					return -1
				} else {
					res += cherry.grid[i][0]
				}
			}
			return res
		} else {
			xMove := cherry.upCount(x-1, y)
			yMove := cherry.upCount(x, y-1)
			if xMove == -1 && yMove == -1 {
				return -1
			} else if xMove > yMove {
				return xMove
			} else {
				return yMove
			}
		}
	}
}
*/

func CherryPickup(grid [][]int) int {
	Cherry := new(farm)
	Cherry.grid = grid
	Cherry.length = len(grid)
	first := Cherry.pickCount(0, 0)
	second := Cherry.pickCount(0, 0)
	return first + second
}
