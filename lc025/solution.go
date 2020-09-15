package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, n int) *ListNode {
	var resu, ptr *ListNode
	resu = &ListNode{
		Next: head,
	}
	ptr = resu
	// include an extra node
	for checkLenLeft(ptr) > n {
		ptr = reverseNextK(ptr, n)
	}
	return resu.Next
}

func checkLenLeft(head *ListNode) int {
	var ptr *ListNode
	ptr = head
	counter := 0
	for ptr != nil {
		counter++
		ptr = ptr.Next
	}
	return counter
}

func reverseNextK(head *ListNode, k int) *ListNode {
	var ptr, last, temp, nextHead *ListNode
	ptr = head
	for k > 0 {
		ptr = ptr.Next
		k--
	}
	nextHead = ptr.Next
	ptr.Next = nil

	ptr = head.Next.Next
	last = ptr.Next
	last.Next = nil
	for temp != nil && temp.Next != nil {
		ptr.Next = last
		last = ptr
		ptr = temp
		temp = temp.Next
	}
	head.Next = ptr
	for ptr != nil {
		ptr = ptr.Next
	}
	ptr.Next = nextHead

	return ptr
}
