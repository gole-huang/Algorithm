package Algo

func compareRoot(first, second *TreeNode) bool {
	if first == nil && second == nil {
		return true
	} else if first == nil || second == nil {
		return false
	}
	if first.Val != second.Val {
		return false
	} else {
		return compareRoot(first.Left, second.Left) && compareRoot(first.Right, second.Right)
	}
}

func dupRoot(node, root *TreeNode, res []*TreeNode) []*TreeNode {
	if node != root {
		if node == nil || root == nil {
			return res
		} else if compareRoot(node, root) {
			match := false
			for _, v := range res {
				if compareRoot(node, v) {
					match = true
					break
				}
			}
			if !match {
				res = append(res, node)
			}
		}
	}
	res = dupRoot(node, root.Left, res)
	res = dupRoot(node, root.Right, res)
	return res
}

func visitTree(node, root *TreeNode, res []*TreeNode) []*TreeNode {
	if node != nil && root != nil {
		if node.Left != nil {
			tmp := node.Left
			node.Left = &TreeNode{500, nil, nil}
			res = dupRoot(tmp, root, res)
			node.Left = tmp
			res = visitTree(node.Left, root, res)
		}
		if node.Right != nil {
			tmp := node.Right
			node.Right = &TreeNode{500, nil, nil}
			res = dupRoot(tmp, root, res)
			node.Right = tmp
			res = visitTree(node.Right, root, res)
		}
	}
	return res
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	check := root
	subTree := false
	for check != nil {
		if check.Left != nil && check.Right != nil {
			subTree = true
			break
		} else if check.Left != nil {
			check = check.Left
		} else if check.Right != nil {
			check = check.Right
		} else {
			break
		}
	}
	if !subTree {
		return nil
	}
	res := visitTree(root, root, make([]*TreeNode, 0))
	return res
}
