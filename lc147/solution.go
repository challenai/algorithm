package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	rsu := &ListNode{}
	ptr1 := head
	ptr2 := rsu
	for ptr1 != nil {
		tempPtr := ptr1
		ptr1 = ptr1.Next
		for tempPtr.Val > ptr2.Val {
			ptr2 = ptr2.Next
		}
		tempPtr.Next = ptr2
		par.Next = tempPtr
	}
}
