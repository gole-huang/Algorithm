package Algo

import "fmt"

/*
type farm struct {
	grid   [][]int
	length int
}
*/

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
			pathX := findPath(x+1, y, cherry)
			if pathX != nil {
				for i := range pathX {
					pathX[i] = append(pathX[i], [2]int{x, y})
				}
			}
		}
		if y != length-1 {
			pathY := findPath(x, y+1, cherry)
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

func CherryPickup(Cherry [][]int) int {
	//length := len(Cherry)
	fmt.Printf("Begin grid: %v\n", Cherry)
	//计算每一条路径中，第一次和第二次采集樱桃的总数
	total := 0
	for _, path := range findPath(0, 0, Cherry) {
		first := 0
		cherry := Cherry
		fmt.Printf("Before grid: %v\n", cherry)
		for i := 0; i < len(path); i++ {
			first += cherry[path[i][0]][path[i][1]]
			cherry[path[i][0]][path[i][1]] = 0
		}
		fmt.Printf("After grid: %v\n", cherry)
		second := pickMost(0, 0, cherry)
		fmt.Printf("first: %v\nSecond: %v\n", first, second)
		if total < first+second {
			total = first + second
		}
	}
	return total
}
