package templates

type ListNode struct {
	Val  int
	Next *ListNode
}

func ll() {
	list := &ListNode{
		Val: 1,
	}
	list.Next = &ListNode{
		Val: 2,
	}
	list.Next.Next = &ListNode{
		Val: 3,
	}
	list.Next.Next.Next = &ListNode{
		Val: 4,
	}
	list.Next.Next.Next = &ListNode{
		Val: 5,
	}
	list.Next.Next.Next.Next = &ListNode{
		Val: 6,
	}

	list2 := &ListNode{
		Val: 5,
	}
	list2.Next = &ListNode{
		Val: 6,
	}
	list2.Next.Next = &ListNode{
		Val: 4,
	}
}
