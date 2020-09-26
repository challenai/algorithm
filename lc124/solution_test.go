package solution

import (
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	root := &TreeNode{
		Val: -10,
	}
	root.Left = &TreeNode{
		Val: 9,
	}
	root.Right = &TreeNode{
		Val: 20,
	}
	root.Right.Left = &TreeNode{
		Val: 15,
	}
	root.Right.Right = &TreeNode{
		Val: 7,
	}
	r := 42
	resu := maxPathSum(root)

	if resu != r {
		t.Fail()
	}
}
