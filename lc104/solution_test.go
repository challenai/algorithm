package solution

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {

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
	r := 4
	resu := maxDepth(root)
	if resu != r {
		t.Fail()
	}
}
