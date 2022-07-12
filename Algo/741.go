package Algo

import "fmt"

type farm struct {
	grid   [][]int
	length int
}

/*
TODO:findPath(x,y) [][][2]int
若cherry(x,y)==-1，则返回nil；
否则若x==y==cherry.length-1，则返回[][][2]int{{{x,y}}}
否则若x==cherry.length-1，则在findPath(x,y+1)的每一个切片后添加(x,y);
	pathY := findPath(x,y+1)
	for i, _ := range pathY {
		pathY[i] = append(pathY[i],{x,y})
	}
否则若y==cherry.length-1，则在findPath(x+1,y)的每一个切片后添加(x,y);
	pathX := findPath(x+1,y)
	for i, _ := range pathX {
		pathX[i] = append(pathX[i],{x,y})
	}
否则添加findPath(x,y+1)，(x,y)和findPath(x+1,y)，(x,y);
	pathX := findPath(x+1,y)
	for i, _ := range pathX {
		pathX[i] = append(pathX[i],{x,y})
	}
	pathY := findPath(x,y+1)
	for i, _ := range pathY {
		pathY[i] = append(pathY[i],{x,y})
	}
	paths = append(pathX, pathY...)
*/
func (cherry *farm) findPath(x, y int) [][][2]int {
	if cherry.grid[x][y] == -1 {
		return nil
	}
	if x == cherry.length-1 && y == cherry.length-1 {
		return [][][2]int{{{x, y}}}
	} else if x == cherry.length-1 {
		pathY := cherry.findPath(x, y+1)
		for i := range pathY {
			pathY[i] = append(pathY[i], [2]int{x, y})
		}
		return pathY
	} else if y == cherry.length-1 {
		pathX := cherry.findPath(x+1, y)
		if pathX != nil {
			for i := range pathX {
				pathX[i] = append(pathX[i], [2]int{x, y})
			}
		}
		return pathX
	} else {
		pathX := cherry.findPath(x+1, y)
		if pathX != nil {
			for i := range pathX {
				pathX[i] = append(pathX[i], [2]int{x, y})
			}
		}
		pathY := cherry.findPath(x, y+1)
		if pathY != nil {
			for i := range pathY {
				pathY[i] = append(pathY[i], [2]int{x, y})
			}
		}
		paths := append(pathX, pathY...)
		return paths
	}
}

func (cherry *farm) pickMost(x, y int) int {
	if cherry.grid[x][y] == -1 {
		return -1
	} else {
		res := cherry.grid[x][y]
		if x == cherry.length-1 && y == cherry.length-1 {
			return res
		} else if x == cherry.length-1 {
			tmpRes := cherry.pickMost(x, y+1)
			if tmpRes != -1 {
				tmpRes += res
			}
			return tmpRes
		} else if y == cherry.length-1 {
			tmpRes := cherry.pickMost(x+1, y)
			if tmpRes != -1 {
				tmpRes += res
			}
			return tmpRes
		} else {
			xMove := cherry.pickMost(x+1, y)
			yMove := cherry.pickMost(x, y+1)
			if xMove == -1 && yMove == -1 {
				fmt.Printf("grid[%d][%d] back to %d\n", x, y, res)
				return -1
			} else if xMove > yMove {
				return res + xMove
			} else {
				return res + yMove
			}
		}
	}
}

func CherryPickup(grid [][]int) int {
	Cherry := new(farm)
	Cherry.grid = grid
	Cherry.length = len(grid)
	var maxPath [][2]int
	first := 0
	for _, path := range Cherry.findPath(0, 0) {
		cnt := 0
		for i := 0; i < len(path); i++ {
			cnt += Cherry.grid[path[i][0]][path[i][1]]
		}
		if first < cnt {
			first = cnt
			maxPath = path
		}
	}
	for _, node := range maxPath {
		Cherry.grid[node[0]][node[1]] = 0
	}
	fmt.Printf("After first time:\n%v\nResult: %v\n", Cherry.grid, first)
	if first == 0 {
		return 0
	}
	second := Cherry.pickMost(0, 0)
	fmt.Printf("After Second time:\n%v\nResult: %v\n", Cherry.grid, second)
	return first + second
}
