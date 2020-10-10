package solution

import (
	"testing"
)

func TestCountNodes(t *testing.T) {
	root := &TreeNode{
		Val: 5,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Left.Left = &TreeNode{
		Val: 1,
	}
	root.Left.Left.Left = &TreeNode{
		Val: 7,
	}
	root.Left.Left.Right = &TreeNode{
		Val: 12,
	}
	root.Left.Right = &TreeNode{
		Val: 4,
	}
	root.Left.Right.Left = &TreeNode{
		Val: 13,
	}
	root.Right = &TreeNode{
		Val: 8,
	}
	root.Right.Left = &TreeNode{
		Val: 3,
	}
	root.Right.Right = &TreeNode{
		Val: 9,
	}
	r := 10
	resu := countNodes(root)

	if resu != r {
		t.Fail()
	}
}
