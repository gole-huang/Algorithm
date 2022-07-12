package Algo

func findPath(x, y int, cherry [][]int) [][][2]int {
	length := len(cherry)
	if cherry[x][y] == -1 {
		return nil
	}
	if x == length-1 && y == length-1 {
		return [][][2]int{{{x, y}}}
	} else {
		var pathX, pathY [][][2]int
		if x != length-1 {
			pathX = findPath(x+1, y, cherry)
			if pathX != nil {
				for i := range pathX {
					pathX[i] = append(pathX[i], [2]int{x, y})
				}
			}
		}
		if y != length-1 {
			pathY = findPath(x, y+1, cherry)
			if pathY != nil {
				for i := range pathY {
					pathY[i] = append(pathY[i], [2]int{x, y})
				}
			}
		}
		paths := append(pathX, pathY...)
		return paths
	}
}

func pickMost(x, y int, cherry [][]int) int {
	length := len(cherry)
	if cherry[x][y] == -1 {
		return -1
	} else {
		res := cherry[x][y]
		if x == length-1 && y == length-1 {
			return res
		} else if x == length-1 {
			tmpRes := pickMost(x, y+1, cherry)
			if tmpRes != -1 {
				tmpRes += res
			}
			return tmpRes
		} else if y == length-1 {
			tmpRes := pickMost(x+1, y, cherry)
			if tmpRes != -1 {
				tmpRes += res
			}
			return tmpRes
		} else {
			xMove := pickMost(x+1, y, cherry)
			yMove := pickMost(x, y+1, cherry)
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

func CherryPickup(cherry [][]int) int {
	total := 0
	for _, path := range findPath(0, 0, cherry) {
		first := 0
		grid := make([][]int, len(cherry))
		for i := 0; i < len(cherry); i++ {
			grid[i] = make([]int, len(cherry[i]))
			copy(grid[i], cherry[i])
		}
		for i := 0; i < len(path); i++ {
			first += grid[path[i][0]][path[i][1]]
			grid[path[i][0]][path[i][1]] = 0
		}
		second := pickMost(0, 0, grid)
		if total < first+second {
			total = first + second
		}
	}
	return total
}
