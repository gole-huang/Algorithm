package Algo

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}
	fast := head
	cnt := 1
	for k != 0 && fast.Next != nil {
		k--
		cnt++
		fast = fast.Next
	}
	if k == 0 {
		slow := head
		for fast.Next != nil {
			fast = fast.Next
			slow = slow.Next
		}
		fast.Next = head
		res := slow.Next
		slow.Next = nil
		return res
	} else {
		k = (k - 1) % cnt
		return RotateRight(head, k)
	}
}
