package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	// I can realize an O(1) sc algorithm, just change the arrow direction.
	// however It's really easy to lead to faults
	st := []*ListNode{}
	counter := 1
	var par, ptr *ListNode
	head = &ListNode{
		Next: head,
	}
	par = head
	ptr = head.Next
	for ptr != nil && counter <= n {
		if counter >= m {
			st = append(st, ptr)
		} else {
			par = ptr
		}
		counter++
		ptr = ptr.Next
	}
	ptr, par = par, ptr
	for len(st) > 0 {
		ptr.Next = st[len(st)-1]
		st = st[:len(st)-1]
		ptr = ptr.Next
	}
	ptr.Next = par
	return head.Next
}
