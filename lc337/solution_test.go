package solution

import (
	"testing"
)

func TestRob(t *testing.T) {
	root := &TreeNode{
		Val: 3,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Left.Right = &TreeNode{
		Val: 3,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	root.Right.Right = &TreeNode{
		Val: 1,
	}
	r1 := 7

	root1 := &TreeNode{
		Val: 3,
	}
	root1.Left = &TreeNode{
		Val: 4,
	}
	root1.Left.Left = &TreeNode{
		Val: 1,
	}
	root.Left.Right = &TreeNode{
		Val: 3,
	}
	root1.Right = &TreeNode{
		Val: 5,
	}
	root1.Right.Right = &TreeNode{
		Val: 1,
	}
	r2 := 9

	rsu1 := rob(root)
	rsu2 := rob(root1)

	if r1 != rsu1 || r2 != rsu2 {
		t.Fail()
	}
}
