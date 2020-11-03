package solution

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 10,
	}
	root.Left = &TreeNode{
		Val: 5,
	}
	root.Left.Left = &TreeNode{
		Val: 3,
	}
	root.Left.Right = &TreeNode{
		Val: 2,
	}
	root.Left.Left.Left = &TreeNode{
		Val: 3,
	}
	root.Left.Left.Right = &TreeNode{
		Val: -2,
	}
	root.Right.Left.Right = &TreeNode{
		Val: 7,
	}
	root.Right = &TreeNode{
		Val: -3,
	}
	root.Right.Right = &TreeNode{
		Val: 11,
	}
	r := 3

	rsu := pathSum(root)
	if r != rsu {
		t.Fail()
	}
}
