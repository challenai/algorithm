package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(num1 *ListNode, num2 *ListNode) *ListNode {
	resu := &ListNode{
		Val:  0,
		Next: nil,
	}

	var ptr1, ptr2, current *ListNode
	current = resu
	ptr1 = num1
	ptr2 = num2
	needPlusOne := 0
	for ptr1 != nil || ptr2 != nil {
		if ptr1 != nil && ptr2 != nil {
			current.Next = &ListNode{
				Val: ptr1.Val + ptr2.Val + needPlusOne,
			}
			ptr1 = ptr1.Next
			ptr2 = ptr2.Next
		} else if ptr1 != nil {
			current.Next = &ListNode{
				Val: ptr1.Val + needPlusOne,
			}
			ptr1 = ptr1.Next
		} else {
			current.Next = &ListNode{
				Val: ptr2.Val + needPlusOne,
			}
			ptr2 = ptr2.Next
		}
		current = current.Next
		if current.Val > 9 {
			current.Val -= 10
			needPlusOne = 1
		} else {
			needPlusOne = 0
		}
	}

	return resu.Next
}
