package solution

import (
	"testing"
)

func TestConnect2(t *testing.T) {
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
	root.Right.Right = &Node{
		Val: 7,
	}
	resu := connect2(root)

	if resu.Left.Left.Next.Next.Val != 7 {
		t.Fail()
	}
}
