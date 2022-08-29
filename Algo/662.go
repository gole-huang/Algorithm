package Algo

func widthOfBinaryTree(root *TreeNode) int {
	lev, maxWidth := 0, 0
	queue := []*TreeNode{root}
	nums := []int{1}
	for len(queue) != 0 {
		lev++
		tmpQ := make([]*TreeNode, 0)
		tmpN := make([]int, 0)
		for k := range queue {
			if queue[k].Left != nil {
				tmpQ = append(tmpQ, queue[k].Left)
				tmpN = append(tmpN, nums[k]*2)
			}
			if queue[k].Right != nil {
				tmpQ = append(tmpQ, queue[k].Right)
				tmpN = append(tmpN, nums[k]*2+1)
			}
		}
		if len(tmpN) > 0 && maxWidth < tmpN[len(tmpN)-1]-tmpN[0] {
			maxWidth = tmpN[len(tmpN)-1] - tmpN[0]
		}
		queue = tmpQ
		nums = tmpN
	}
	return maxWidth + 1
}
