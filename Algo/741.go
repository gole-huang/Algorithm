package Algo

func findPath(x, y int, cherry [][]int) []map[[2]int]int {
	length := len(cherry)
	if cherry[x][y] == -1 {
		return nil
	}
	if x == length-1 && y == length-1 {
		pathMap := make([]map[[2]int]int, 1)
		pathMap[0] = make(map[[2]int]int, 1)
		pathMap[0][[2]int{x, y}] = cherry[x][y]
		return pathMap
	} else {
		var pathX, pathY []map[[2]int]int
		if x != length-1 {
			pathX = findPath(x+1, y, cherry)
			if pathX != nil {
				for i := range pathX {
					pathX[i][[2]int{x, y}] = cherry[x][y]
				}
			}
		}
		if y != length-1 {
			pathY = findPath(x, y+1, cherry)
			if pathY != nil {
				for i := range pathY {
					pathY[i][[2]int{x, y}] = cherry[x][y]
				}
			}
		}
		paths := append(pathX, pathY...)
		return paths
	}
}

func pickMost(x, y int, cherry [][]int, pickMap map[[2]int]int) int {
	length := len(cherry)
	if cherry[x][y] == -1 {
		return -1
	} else {
		res := cherry[x][y] - pickMap[[2]int{x, y}]
		if x == length-1 && y == length-1 {
			return res
		} else {
			var xMove, yMove int
			if x != length-1 {
				xMove = pickMost(x+1, y, cherry, pickMap)
			}
			if y != length-1 {
				yMove = pickMost(x, y+1, cherry, pickMap)
			}
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
	paths := findPath(0, 0, cherry)
	for _, path := range paths {
		first := 0
		grid := make([][]int, len(cherry))
		for i := 0; i < len(cherry); i++ {
			grid[i] = make([]int, len(cherry[i]))
			copy(grid[i], cherry[i])
		}
		/*
			for i := 0; i < len(path); i++ {
				first += grid[path[i][0]][path[i][1]]
				grid[path[i][0]][path[i][1]] = 0
			}
		*/
		for p := range path {
			first += cherry[p[0]][p[1]]
			cherry[p[0]][p[1]] = 0
		}
		second := pickMost(0, 0, grid, path)
		if total < first+second {
			total = first + second
		}
	}
	return total
}
