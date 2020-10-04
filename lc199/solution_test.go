package solution

import (
	"testing"
)

func TestX(t *testing.T) {
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
	root.Right.Left.Right.Left = &TreeNode{
		Val: 6,
	}
	root.Right.Right = &TreeNode{
		Val: 9,
	}
	r := []int{5, 8, 9, 7, 6}
	resu := rightSideView(root)

	if len(r) != len(resu) {
		t.Fail()
	}
	for i := 0; i < len(r); i++ {
		if resu[i] != r[i] {
			t.Fail()
		}
	}
}
