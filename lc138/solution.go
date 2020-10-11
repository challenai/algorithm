package solution

type Node struct {
	Val          int
	Next, Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	hash := map[*Node]*Node{}
	var ptr0, ptr1, head1 *Node
	ptr0 = head.Next
	head1 = &Node{
		Val: head.Val,
	}
	hash[head] = head1
	ptr1 = head1
	for ptr0 != nil {
		ptr1.Next = &Node{
			Val: ptr0.Val,
		}
		hash[ptr0] = ptr1.Next
		ptr1 = ptr1.Next
		ptr0 = ptr0.Next
	}
	ptr0 = head
	ptr1 = head1
	for ptr0 != nil {
		if ptr0.Random != nil {
			ptr1.Random = hash[ptr0.Random]
		}
		ptr0 = ptr0.Next
		ptr1 = ptr1.Next
	}
	return head1
}
