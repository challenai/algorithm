package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicate2(head *ListNode) *ListNode {
	var ptr, before *ListNode
	head = &ListNode{
		Val:  0,
		Next: head,
	}
	ptr = head.Next.Next
	before = head
	markASDelete := false
	for ptr != nil {
		if ptr.Val == before.Next.Val {
			markASDelete = true
		} else {
			if markASDelete {
				before.Next = ptr
			} else {
				before = before.Next
			}
			markASDelete = false
		}
		ptr = ptr.Next
	}

	ptr = head.Next
	// for ptr != nil {
	// 	print(ptr.Val, " ")
	// 	ptr = ptr.Next
	// }
	// println()

	return head.Next
}
