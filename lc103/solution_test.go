package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	root := &TreeNode{
		Val: 3,
	}
	root.Left = &TreeNode{
		Val: 9,
	}
	root.Right = &TreeNode{
		Val: 20,
	}
	root.Right.Left = &TreeNode{
		Val: 15,
	}
	root.Right.Right = &TreeNode{
		Val: 7,
	}
	r := [][]int{
		{3},
		{20, 9},
		{15, 7},
	}
	resu := zigzagLevelOrder(root)
	for i := 0; i < len(r); i++ {
		if len(r[i]) != len(resu[i]) {
			t.Fail()
			break
		}
	}
}
