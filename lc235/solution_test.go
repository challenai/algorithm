package solution

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
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
		Val: 12,
	}
	root.Right.Left = &TreeNode{
		Val: 8,
	}
	root.Right.Left.Left = &TreeNode{
		Val: 7,
	}
	root.Right.Left.Right = &TreeNode{
		Val: 9,
	}
	root.Right.Right = &TreeNode{
		Val: 13,
	}
	root.Right.Right.Right = &TreeNode{
		Val: 14,
	}

	p := root.Right.Left.Right
	q := root.Right.Right.Right
	r := root.Right

	resu := lowestCommonAncestor(root, p, q)

	if resu != r {
		t.Fail()
	}
}
