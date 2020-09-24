package solution

import (
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 2,
	}
	root.Left.Left = &TreeNode{
		Val: 3,
	}
	root.Left.Right = &TreeNode{
		Val: 4,
	}
	root.Right.Left = &TreeNode{
		Val: 4,
	}
	root.Right.Right = &TreeNode{
		Val: 3,
	}

	resu := isSymmetric(root)
	if !resu {
		t.Fail()
	}
}
