package Algo

func addNode(value int, left, right []int) (root *TreeNode) {
	root = new(TreeNode)
	root.Val = value
	if len(left) == 0 {
		root.Left = nil
	} else {
		max, l := 0, 0
		for k := range left {
			if max < left[k] {
				max = left[k]
				l = k
			}
		}
		root.Left = addNode(left[l], left[:l], left[l+1:])
	}
	if len(right) == 0 {
		root.Right = nil
	} else {
		max, r := 0, 0
		for k := range right {
			if max < right[k] {
				max = right[k]
				r = k
			}
		}
		root.Right = addNode(right[r], right[:r], right[r+1:])
	}
	return root
}

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	max, center := 0, 0
	for k := range nums {
		if max < nums[k] {
			max = nums[k]
			center = k
		}
	}
	return addNode(nums[center], nums[:center], nums[center+1:])
}
