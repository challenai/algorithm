package solution

import (
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 5,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Left.Left = &TreeNode{
		Val: 1,
	}
	root.Left.Right = &TreeNode{
		Val: 3,
	}
	root.Right = &TreeNode{
		Val: 8,
	}
	root.Right.Left = &TreeNode{
		Val: 6,
	}
	root.Right.Left.Right = &TreeNode{
		Val: 7,
	}
	root.Right.Right = &TreeNode{
		Val: 9,
	}
	r := []int{1, 3, 2, 7, 6, 9, 8, 5}

	resu := postorderTraversal(root)

	if len(r) != len(r) {
		t.Fail()
	}
	for i := 0; i < len(resu); i++ {
		if r[i] != resu[i] {
			t.Fail()
		}
	}
}
