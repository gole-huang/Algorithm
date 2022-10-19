package algo

func NumComponents(head *ListNode, nums []int) int {
	cnt, node, newApponent := 0, head, false
	for node != nil && len(nums) > 0 {
		d := len(nums) - 1
		for ; d >= 0; d-- {
			if nums[d] == node.Val {
				if !newApponent {
					cnt++
				}
				newApponent = true
				nums = append(nums[:d], nums[d+1:]...)
				break
			}
		}
		if d == -1 {
			newApponent = false
		}
		node = node.Next
	}
	return cnt
}
