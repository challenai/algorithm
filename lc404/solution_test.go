package solution

import (
	"testing"
)

func TestSumOfLeftLeaves(t *testing.T) {
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
		Val: 4,
	}
	root.Left.Right.Left = &TreeNode{
		Val: 3,
	}
	root.Right = &TreeNode{
		Val: 9,
	}
	root.Right.Left = &TreeNode{
		Val: 7,
	}
	root.Right.Left.Left = &TreeNode{
		Val: 6,
	}
	root.Right.Left.Right = &TreeNode{
		Val: 8,
	}
	root.Right.Right = &TreeNode{
		Val: 11,
	}
	r := 10

	sum := sumOfLeftLeaves(root)
	println(sum)

	if sum != r {
		t.Fail()
	}
}
