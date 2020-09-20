package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotate(head *ListNode, k int) *ListNode {
	var resu *ListNode
	counter := 0
	var ptr *ListNode
	ptr = head
	for ptr != nil {
		counter++
		ptr = ptr.Next
	}
	k %= counter
	if k == 0 {
		return head
	}
	counter -= k
	counter--
	ptr = head
	for ptr != nil && counter > 0 {
		counter--
		ptr = ptr.Next
	}
	resu = ptr.Next
	ptr.Next = nil
	ptr = resu
	for ptr != nil && ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = head
	return resu
}
