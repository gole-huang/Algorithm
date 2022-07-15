package Algo

import "fmt"

func firstPath(x, y int, cherry [][]int) [][][2]int {
	//fmt.Printf("(%v,%v\n", x, y)
	length := len(cherry)
	if cherry[x][y] == -1 {
		return nil
	}
	if x == length-1 && y == length-1 {
		return [][][2]int{{{x, y}}}
	} else {
		var pathX, pathY [][][2]int
		if x != length-1 {
			pathX = firstPath(x+1, y, cherry)
			if pathX != nil {
				for i := 0; i < len(pathX); i++ {
					pathX[i] = append(pathX[i], [2]int{x, y})
				}
			}
		}
		if y != length-1 {
			pathY = firstPath(x, y+1, cherry)
			if pathY != nil {
				for i := 0; i < len(pathY); i++ {
					pathY[i] = append(pathY[i], [2]int{x, y})
				}
			}
		}
		paths := append(pathX, pathY...)
		return paths
	}
}

func secondPath(x, y, target int, cherry [][]int) int {
	length := len(cherry)
	if cherry[x][y] == -1 || target >= (length)*2-1-x-y {
		return -1
	} else {
		res := cherry[x][y]
		target -= res
		if x == length-1 && y == length-1 {
			return res
		} else {
			//！！！！！必须要给予xMove/yMove初始值-1！！！！！
			xMove, yMove := -1, -1
			if x != length-1 {
				xMove = secondPath(x+1, y, target, cherry)
			}
			if y != length-1 {
				yMove = secondPath(x, y+1, target, cherry)
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
	fmt.Printf("Start\n")
	paths := firstPath(0, 0, cherry)
	fmt.Printf("Find all paths\n")
	for p := 0; p < len(paths); p++ {
		first := 0
		grid := make([][]int, len(cherry))
		for i := 0; i < len(cherry); i++ {
			grid[i] = make([]int, len(cherry[i]))
			copy(grid[i], cherry[i])
		}
		for i := 0; i < len(paths[p]); i++ {
			first += grid[paths[p][i][0]][paths[p][i][1]]
			grid[paths[p][i][0]][paths[p][i][1]] = 0
		}
		if first < total>>1 {
			continue
		}
		second := secondPath(0, 0, total-first, grid)
		if total < first+second {
			total = first + second
		}
	}
	return total
}
