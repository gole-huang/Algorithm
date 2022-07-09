package Algo

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := new(ListNode)
	node.Next = head
	newHead := head.Next
	for node.Next != nil && node.Next.Next != nil {
		tmp := node.Next
		tail := node.Next.Next.Next
		node.Next = node.Next.Next
		node.Next.Next = tmp
		node.Next.Next.Next = tail
		node = node.Next.Next
	}
	return newHead
}
