package solution

import (
	"testing"
)

func TestHasPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 5,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Left.Left = &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val: 8,
	}
	root.Right.Left = &TreeNode{
		Val: 3,
	}
	root.Right.Left.Right = &TreeNode{
		Val: 7,
	}
	root.Right.Right = &TreeNode{
		Val: 9,
	}
	resu1 := hasPathSum(root, 23)
	resu2 := hasPathSum(root, 51)

	if !resu1 || resu2 {
		t.Fail()
	}
}
