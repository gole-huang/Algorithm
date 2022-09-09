package algo

import "strconv"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getHeight(t *TreeNode) int {
	if t.Left == nil && t.Right == nil {
		return 0
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
	width := 2<<height - 1
	res := [][3]int{{0, (width - 1) >> 1, t.Val}}
	layer, cur := 0, 0
	nodes := []*TreeNode{t}
	for len(nodes) != 0 {
		layer++
		tmpNodes := make([]*TreeNode, 0)
		for _, v := range nodes {
			if v.Left != nil {
				tmpNodes = append(tmpNodes, v.Left)
				if height == layer {
					res = append(res, [3]int{layer, res[cur][1] - 1, v.Left.Val})
				} else {
					res = append(res, [3]int{layer, res[cur][1] - 2<<(height-layer-1), v.Left.Val})
				}
			}
			if v.Right != nil {
				tmpNodes = append(tmpNodes, v.Right)
				if height == layer {
					res = append(res, [3]int{layer, res[cur][1] + 1, v.Right.Val})
				} else {
					res = append(res, [3]int{layer, res[cur][1] + 2<<(height-layer-1), v.Right.Val})
				}
			}
			cur++
		}
		nodes = tmpNodes
	}
	return res
}

func PrintTree(root *TreeNode) [][]string {
	height := getHeight(root) // height=树的高度+1
	res := make([][]string, height+1)
	width := 2<<height - 1
	for k := range res {
		res[k] = make([]string, width)
	}
	for _, v := range bfs(root) {
		res[v[0]][v[1]] = strconv.Itoa(v[2])
	}
	return res
}
