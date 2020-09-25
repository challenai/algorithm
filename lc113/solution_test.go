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
		Val: 6,
	}
	root.Right.Right = &TreeNode{
		Val: 9,
	}
	resu := pathSum(root, 22)

	if len(resu) != 2 {
		t.Fail()
	}
}
