package algo

/*
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
*/
func DeepestLeavesSum(root *TreeNode) int {
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		tmpNodes := make([]*TreeNode, 0)
		for _, v := range nodes {
			if v.Left != nil {
				tmpNodes = append(tmpNodes, v.Left)
			}
			if v.Right != nil {
				tmpNodes = append(tmpNodes, v.Right)
			}
		}
		if len(tmpNodes) != 0 {
			nodes = tmpNodes
		} else {
			res := 0
			for _, v := range nodes {
				res += v.Val
			}
			return res
		}
	}
	return root.Val
}
