package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(n *ListNode) {
	if n == nil {
		return
	}
	if n.Next == nil {
		n = nil
		return
	}
	n.Val = n.Next.Val
	n.Next = n.Next.Next
}
