package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	r1 := 25
	root2 := &TreeNode{
		Val: 4,
	}
	root2.Left = &TreeNode{
		Val: 9,
	}
	root2.Left.Left = &TreeNode{
		Val: 5,
	}
	root2.Left.Right = &TreeNode{
		Val: 1,
	}
	root2.Right = &TreeNode{
		Val: 0,
	}
	r2 := 1026
	resu1 := sumNumbers(root)
	resu2 := sumNumbers(root2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
