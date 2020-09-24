package solution

import (
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {
	root := &TreeNode{
		Val: 5,
	}
	root.Left = &TreeNode{
		Val: 2,
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
	r := [][]int{
		{7},
		{3, 9},
		{2, 8},
		{5},
	}
	resu := levelOrderBottom(root)
	for i := 0; i < len(r); i++ {
		if len(r[i]) != len(resu[i]) {
			t.Fail()
			break
		}
	}
}
