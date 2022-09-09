package algo

import (
	"fmt"
)

type Cherry struct {
	Grid [][][2]bool
}

// 把果园从int64转换为2*bool，压缩内存
func (grid *Cherry) transform(cherry [][]int) {
	grid.Grid = make([][][2]bool, len(cherry))
	for i := 0; i < len(cherry); i++ {
		grid.Grid[i] = make([][2]bool, len(cherry[i]))
		for j := range cherry[i] {
			if cherry[i][j] == -1 {
				grid.Grid[i][j] = [2]bool{true, false}
			} else if cherry[i][j] == 1 {
				grid.Grid[i][j] = [2]bool{false, true}
			} else {
				grid.Grid[i][j] = [2]bool{false, false}
			}
		}
	}
}

func (grid *Cherry) copyFrom(cherry *Cherry) {
	grid.Grid = make([][][2]bool, len(cherry.Grid))
	for i := 0; i < len(cherry.Grid); i++ {
		grid.Grid[i] = make([][2]bool, len(cherry.Grid[i]))
		for j := range cherry.Grid[i] {
			grid.Grid[i][j] = cherry.Grid[i][j]
		}
	}
}

func firstPath(x, y int8, cherry *Cherry) [][][2]int8 {
	length := int8(len(cherry.Grid))
	var pathX, pathY [][][2]int8
	if x == length-1 && y == length-1 {
		return [][][2]int8{{{x, y}}}
	} else if x == length-1 {
		/*pathY = make([][][2]int8, 1)
		pathY[0] = make([][2]int8, 0)
		for i := length - 1; i >= y; i-- {
			if cherry.Grid[x][i][0] {
				return nil
			} else {
				pathY[0] = append(pathY[0], [2]int8{x, i})
			}
		}
		return pathY*/
		for i := length - 1; i >= y; i-- {
			if cherry.Grid[x][i][0] {
				return nil
			}
		}
		return [][][2]int8{{{x, y}}}
	} else if y == length-1 {
		/*pathX = make([][][2]int8, 1)
		pathX[0] = make([][2]int8, 0)
		for i := length - 1; i >= x; i-- {
			if cherry.Grid[i][y][0] {
				return nil
			} else {
				pathX[0] = append(pathX[0], [2]int8{i, y})
			}
		}
		return pathX*/
		for i := length - 1; i >= x; i-- {
			if cherry.Grid[i][y][0] {
				return nil
			}
		}
		return [][][2]int8{{{x, y}}}
	} else {
		//对于某一个节点，需要到达右下节点可通过2条路经，除非先右后下比先下后右大，否则都先下后右
		/*当且仅当满足下述条件，需要添加下方路径集合：
		**下方单元格不为-1，且
		**没有左下方，或
		***本单元格为1或左下方单元格为0*/
		if !cherry.Grid[x+1][y][0] { //添加下方路径
			pathX = firstPath(x+1, y, cherry)
			for i := 0; i < len(pathX); i++ {
				if !cherry.Grid[x+1][y][1] && cherry.Grid[x][y+1][1] && (pathX[i] == nil || pathX[i][len(pathX[i])-1] == [2]int8{x + 1, y + 1}) { //当且仅当无左下节点，或者本节点为1或左下节点为0
					pathX[i] = nil
				} else if pathX[i] != nil {
					pathX[i] = append(pathX[i], [2]int8{x, y})
				}
			}
		}
		/*当且仅当满足下述条件，需要添加右方路径集合：
		**右方单元格不为-1，且
		**没有右上方，或
		***本单元格为1且右上方单元格为0*/
		if !cherry.Grid[x][y+1][0] { //添加右方路径
			pathY = firstPath(x, y+1, cherry)
			for i := 0; i < len(pathY); i++ {
				if (cherry.Grid[x+1][y][1] || !cherry.Grid[x][y+1][1]) && (pathY[i] == nil || pathY[i][len(pathY[i])-1] == [2]int8{x + 1, y + 1}) { //当且仅当无右上节点，或者本节点为1且右上节点为0
					pathY[i] = nil
				} else if pathY[i] != nil {
					pathY[i] = append(pathY[i], [2]int8{x, y})
				}
			}
		}
		paths := append(pathX, pathY...)
		/*for _, p := range paths {
			fmt.Printf("Each path: %v\n", p)
		}*/
		return paths
	}
}

func secondPath(x, y int8, target int, cherry *Cherry) int {
	length := int8(len(cherry.Grid))
	if cherry.Grid[x][y][0] || target >= int((length)*2-1-x+y) {
		return -1
	} else {
		var res int
		if cherry.Grid[x][y][1] {
			res = 1
		} else {
			res = 0
		}
		target -= res
		if x == length-1 && y == length-1 {
			return res
		} else {
			//！！！！！必须要给予xMove/yMove初始值-1！！！！！
			xMove, yMove := -1, -1
			if x != length-1 && !cherry.Grid[x+1][y][0] {
				xMove = secondPath(x+1, y, target, cherry)
			}
			if y != length-1 && !cherry.Grid[x][y+1][0] {
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

func CherryPickup(cherryInt [][]int) int {
	total := 0
	fmt.Printf("Start\n")
	cherry := new(Cherry)
	cherry.transform(cherryInt)
	pathLength := len(cherry.Grid)<<1 - 1
	paths := firstPath(0, 0, cherry)
	fmt.Printf("Find all paths\n")
	for p := 0; p < len(paths); p++ {
		first := 0
		grid := new(Cherry)
		grid.copyFrom(cherry)
		//需要检查每一条路径的第一位是否为(length-1,length-1)，否则补全到最后路径
		if len(paths[p]) < pathLength {
			if int(paths[p][0][0]) < len(cherry.Grid) { //若第一个元素的x不足，则沿着(x,len(cherry.Grid))补全
				for i := len(cherry.Grid) - 1; i >= int(paths[p][0][0]); i-- {
					if grid.Grid[i][len(cherry.Grid)-1][1] {
						first++
					}
				}
			} else { //否则沿着(len(cherry.Grid),y)补全
				for i := len(cherry.Grid) - 1; i >= int(paths[p][0][1]); i-- {
					if grid.Grid[len(cherry.Grid)-1][i][1] {
						first++
					}
				}
			}
		}
		for i := 0; i < len(paths[p]); i++ {
			if grid.Grid[paths[p][i][0]][paths[p][i][1]][1] {
				first++
			}
			grid.Grid[paths[p][i][0]][paths[p][i][1]] = [2]bool{false, false}
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
