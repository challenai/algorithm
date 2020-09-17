package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var ptr1, ptr2, ptr3, resu *ListNode
	resu = &ListNode{}
	ptr1 = l1
	ptr2 = l2
	ptr3 = resu
	for ptr1 != nil || ptr2 != nil {
		if ptr1 != nil && ptr2 != nil {
			if ptr1.Val < ptr2.Val {
				ptr3.Next = &ListNode{
					Val: ptr1.Val,
				}
				ptr1 = ptr1.Next
			} else {
				ptr3.Next = &ListNode{
					Val: ptr2.Val,
				}
				ptr2 = ptr2.Next
			}
		} else if ptr1 != nil {
			ptr3.Next = &ListNode{
				Val: ptr1.Val,
			}
			ptr1 = ptr1.Next
		} else {
			ptr3.Next = &ListNode{
				Val: ptr2.Val,
			}
			ptr2 = ptr2.Next
		}
	}

	return resu.Next
}
