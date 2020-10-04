package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, v int) *ListNode {
	var ptr, par *ListNode
	resu := &ListNode{
		Next: head,
	}
	par = resu
	ptr = head
	for ptr != nil {
		if ptr.Val == v {
			par.Next = par.Next.Next
			ptr = ptr.Next
			continue
		}
		par = ptr
		ptr = ptr.Next
	}
	return resu.Next
}
