package Algo

import (
	"math"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getHeight(t *TreeNode) int {
	if t.Left == nil && t.Right == nil {
		return 1
	} else if t.Right == nil {
		return 1 + getHeight(t.Left)
	} else if t.Left == nil {
		return 1 + getHeight(t.Right)
	} else {
		return max(1+getHeight(t.Left), 1+getHeight(t.Right))
	}
}

func bfs(t *TreeNode) [][3]int {
	height := getHeight(t)
	res := [][3]int{{0, (height - 1) / 2, t.Val}}
	layer, cur := 0, 0
	nodes := []*TreeNode{t}
	for nodes != nil {
		tmpNodes := make([]*TreeNode, 0)
		for _, v := range nodes {
			if v.Left != nil {
				tmpNodes = append(tmpNodes, v.Left)
				res = append(res, [3]int{layer, res[cur][1] - int(math.Pow(2, float64(height-1-layer))), v.Left.Val})
			}
			if v.Right != nil {
				tmpNodes = append(tmpNodes, v.Right)
				res = append(res, [3]int{layer, res[cur][1] + int(math.Pow(2, float64(height-1-layer))), v.Right.Val})
			}
			cur++
		}
		nodes = tmpNodes
		layer++
	}
	return res
}

//PrintTree
/*
树的高度为height，行数m = height+1。
矩阵的列数n = 2^(height+1)-1。
根节点需要放置在顶行的正中间，对应位置为res[0][(n-1)/2]。
对于放置在矩阵中的每个节点，设对应位置为res[r][c]，将其左子节点放置在res[r+1][c-2^(height-r-1)]右子节点放置在res[r+1][c+2^(height-(r+1))]。
任意空单元格都应该包含空字符串 "" 。
*/
func PrintTree(root *TreeNode) [][]string {
	height := getHeight(root) //height=树的高度+1
	res := make([][]string, height)
	width := int(math.Pow(2, float64(height))) - 1
	for k := range res {
		res[k] = make([]string, width)
	}
	for _, v := range bfs(root) {
		res[v[0]][v[1]] = strconv.Itoa(v[2])
	}
	return res
}
