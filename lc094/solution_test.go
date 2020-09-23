package solution

import (
	"testing"
)

func TestInorderTraversal(t *testing.T) {

	root := &TreeNode{
		Val: 5,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 8,
	}
	root.Left.Left = &TreeNode{
		Val: 1,
	}
	root.Left.Right = &TreeNode{
		Val: 3,
	}
	root.Left.Right.Right = &TreeNode{
		Val: 4,
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
	resu := inorderTraversal2(root)
	if len(resu) != 9 {
		t.Fail()
	}
	for i := 1; i < len(resu); i++ {
		if resu[i] != resu[i-1]+1 {
			t.Fail()
		}
	}
}
