package solution

import (
	"testing"
)

func TestConnect(t *testing.T) {
	root := &Node{
		Val: 1,
	}
	root.Left = &Node{
		Val: 2,
	}
	root.Right = &Node{
		Val: 3,
	}
	root.Left.Left = &Node{
		Val: 4,
	}
	root.Left.Right = &Node{
		Val: 5,
	}
	root.Right.Left = &Node{
		Val: 6,
	}
	root.Right.Right = &Node{
		Val: 7,
	}
	resu := connect(root)

	if resu.Left.Next.Left.Next.Val != 7 {
		t.Fail()
	}
}
