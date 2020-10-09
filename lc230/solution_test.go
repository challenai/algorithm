package solution

import (
	"testing"
)

func TestKthSmallest(t *testing.T) {
	root := &TreeNode{
		Val: 5,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Left.Right = &TreeNode{
		Val: 3,
	}
	root.Left.Left = &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val: 8,
	}
	root.Right.Left = &TreeNode{
		Val: 7,
	}
	root.Right.Right = &TreeNode{
		Val: 9,
	}
	k := 5
	r := 7
	resu := kthSmallest(root, k)

	if resu != r {
		t.Fail()
	}
}
