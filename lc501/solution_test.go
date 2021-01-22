package solution

import (
	"testing"
)

func TestFindMode(t *testing.T) {

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
		Val: 2,
	}
	root.Left.Right.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 8,
	}
	root.Right.Left = &TreeNode{
		Val: 6,
	}
	root.Right.Left.Right = &TreeNode{
		Val: 6,
	}
	root.Right.Right = &TreeNode{
		Val: 9,
	}
	r := 2

	rsu := findMode(root)
	if rsu != r {
		t.Fail()
	}
}
