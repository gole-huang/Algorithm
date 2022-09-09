package algo

func InsertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if val > root.Val {
		head := new(TreeNode)
		head.Val = val
		head.Left = root
		return head
	}
	cur := root
	node := new(TreeNode)
	node.Val = val
	for cur.Right != nil && cur.Right.Val > val {
		cur = cur.Right
	}
	if cur.Right != nil {
		node.Left = cur.Right
	}
	cur.Right = node
	return root
}
