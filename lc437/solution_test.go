package solution

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 8,
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
	root.Left.Right.Right = &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val: -3,
	}
	root.Right.Right = &TreeNode{
		Val: 11,
	}
	sum := 8
	r := 4

	rsu := pathSum(root, sum)
	if r != rsu {
		t.Fail()
	}
}
