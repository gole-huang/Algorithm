package algo

func Partition(head *ListNode, x int) *ListNode {
    if head!=nil && head.Next!=nil {
	    front, tail:=new(ListNode),new(ListNode)
	    front.Next=head
        head=nil
	    rear:=tail
	    for front.Next != nil {
		    if front.Next.Val >= x {
			    rear.Next=front.Next
			    rear=rear.Next
			    front.Next=rear.Next
			    rear.Next=nil
		    } else {
                if head==nil{
                    head=front.Next
                }
		        front=front.Next
            }
	    }
	    front.Next=tail.Next
        if head==nil {
            head=tail.Next
        }
    }
	return head
}