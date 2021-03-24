package solution

import (
	"testing"
)

func TestLargestValue(t *testing.T) {
	root := &TreeNode{
		Val: 5,
	}
	root.Left = &TreeNode{
		Val: -2,
	}
	root.Left.Left = &TreeNode{
		Val: 1,
	}
	root.Left.Right = &TreeNode{
		Val: -3,
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
		Val: -9,
	}
	r := []int{5, 8, 6, 7}

	rsu := LargestValue(root)

	if len(r) != len(rsu) {
		t.Fail()
	}
	for i := 0; i < len(r); i++ {
		if r[i] != rsu[i] {
			t.Fail()
		}
	}
}
