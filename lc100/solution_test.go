package solution

import (
	"testing"
)

func TestIsSameTree(t *testing.T) {

	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Left.Right = &TreeNode{
		Val: 3,
	}
	root.Right = &TreeNode{
		Val: 4,
	}
	root1 := &TreeNode{
		Val: 1,
	}
	root1.Left = &TreeNode{
		Val: 2,
	}
	root1.Left.Right = &TreeNode{
		Val: 3,
	}
	root1.Right = &TreeNode{
		Val: 4,
	}

	resu := isSameTree(root, root1)
	if !resu {
		t.Fail()
	}
}
