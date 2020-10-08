package solution

import (
	"testing"
)

func TestInvertTree(t *testing.T) {
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
		Val: 7,
	}
	root.Right.Left = &TreeNode{
		Val: 6,
	}
	root.Right.Left.Right = &TreeNode{
		Val: 111,
	}
	root.Right.Right = &TreeNode{
		Val: 9,
	}
	resu := revertTree(root)

	if resu.Left.Val != 7 || resu.Left.Left.Val != 9 {
		t.Fail()
	}
}
