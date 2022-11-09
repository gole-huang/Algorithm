package algo

import "fmt"

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

func miner(a, b [2]int) int {
	row, col := 0, 0
	if a[0] > b[0] {
		row = a[0] - b[0]
	} else {
		row = b[0] - a[0]
	}
	if a[1] > b[1] {
		col = a[1] - b[1]
	} else {
		col = b[1] - a[1]
	}
	return row + col
}

// func buildIsland(grid [][]int, from Direction, x, y int, edge [][2]int, island map[[2]int]bool) ([][2]int,map[[2]int]bool) {
func buildIsland(grid [][]int, from Direction, x, y int, edge [][2]int, island map[[2]int]bool) {
	isEdge := false
	if x > 0 && from != Left {
		if grid[x-1][y] == 0 {
			isEdge = true
		} else {
			//edge,island = buildIsland(grid, Right, x-1, y, edge, island)
			buildIsland(grid, Right, x-1, y, edge, island)
		}
	}
	if x < len(grid)-1 && from != Right {
		if grid[x+1][y] == 0 {
			isEdge = true
		} else {
			//edge,island = buildIsland(grid, Left, x+1, y, edge, island)
			buildIsland(grid, Left, x+1, y, edge, island)
		}
	}
	if y > 0 && from != Up {
		if grid[x][y-1] == 0 {
			isEdge = true
		} else {
			//edge,island = buildIsland(grid, Down, x, y-1, edge, island)
			buildIsland(grid, Down, x, y-1, edge, island)
		}
	}
	if y < len(grid[x])-1 && from != Down {
		if grid[x][y+1] == 0 {
			isEdge = true
		} else {
			//edge,island = buildIsland(grid, Up, x, y+1, edge, island)
			buildIsland(grid, Up, x, y+1, edge, island)
		}
	}
	if isEdge {
		edge = append(edge, [2]int{x, y})
	}
	island[[2]int{x, y}] = true
	//return edge,island
}

func shortestBridge(grid [][]int) int {
	one, two := make([][2]int, 0), make([][2]int, 0)
	tmp := make(map[[2]int]bool)
	for i := range grid {
		findOut := false
		for j := range grid[i] {
			if grid[i][j] == 1 {
				findOut = true
				if j > 0 {
					one, tmp = buildIsland(grid, Left, i, j, one, tmp)
				} else {
					one, tmp = buildIsland(grid, Up, i, j, one, tmp)
				}
				break
			}
		}
		if findOut {
			break
		}
	}
	fmt.Printf("One: %v\t%v\n", one, tmp)
	for i := range grid {
		for j := range grid[i] {
			if tmp[[2]int{i, j}] {
				continue
			}
			if grid[i][j] == 1 {
				if j > 0 {
					two, tmp = buildIsland(grid, Left, i, j, two, tmp)
				} else {
					two, tmp = buildIsland(grid, Up, i, j, two, tmp)
				}
			}
		}
	}
	fmt.Printf("Two: %v\t%v\n", two, tmp)
	min := 1 << 31
	for _, o := range one {
		for _, t := range two {
			if min > miner(o, t) {
				min = miner(o, t)
			}
		}
	}
	return min - 1
}
