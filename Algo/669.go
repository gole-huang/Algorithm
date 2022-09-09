package algo

func TrimBST(root *TreeNode, low int, high int) *TreeNode {
	if root.Left != nil {
		root.Left = TrimBST(root.Left, low, high)
	}
	if root.Val < low {
		if root.Right == nil {
			return nil
		} else {
			root = TrimBST(root.Right, low, high)
		}
	} else if root.Val > high {
		if root.Left == nil {
			return nil
		} else {
			root = TrimBST(root.Left, low, high)
		}
	}
	if root == nil {
		return nil
	}
	if root.Right != nil {
		root.Right = TrimBST(root.Right, low, high)
	}
	return root
}
