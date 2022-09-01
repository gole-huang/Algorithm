package Algo

func triComp(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}

func singlePath(node *TreeNode) int {
	left, right := 0, 0
	if node.Left != nil && node.Left.Val == node.Val {
		left = singlePath(node.Left) + 1
	}
	if node.Right != nil && node.Right.Val == node.Val {
		right = singlePath(node.Right) + 1
	}
	if left < right {
		return right
	} else {
		return left
	}
}

func LongestUnivaluePath(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	left, right := 0, 0
	if root.Left != nil && root.Left.Val == root.Val {
		left = singlePath(root.Left) + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		right = singlePath(root.Right) + 1
	}
	return triComp(left+right, LongestUnivaluePath(root.Left), LongestUnivaluePath(root.Right))
}
