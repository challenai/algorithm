package solution

import (
	"testing"
)

func TestSortedListToBST(t *testing.T) {

	head := &ListNode{
		Val: -10,
	}
	head.Next = &ListNode{
		Val: -3,
	}
	head.Next.Next = &ListNode{
		Val: 0,
	}
	head.Next.Next.Next = &ListNode{
		Val: 5,
	}
	head.Next.Next.Next.Next = &ListNode{
		Val: 9,
	}

	resu := sortedListToBST(head)
	inorderTraversal(resu)

	if resu.Val != 0 {
		t.Fail()
	}
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	print(root.Val, " ")
	inorderTraversal(root.Right)
}
