package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// write down all the pointer in a graph is a great idea
	var temp *ListNode
	par := head
	ptr := head.Next
	par.Next = nil
	for ptr != nil {
		temp = ptr.Next
		ptr.Next = par
		par = ptr
		ptr = temp
	}
	return par
}
