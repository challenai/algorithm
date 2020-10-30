package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	root := &Node{
		Val: 1,
	}
	root.Right = &Node{
		Val: 2,
	}
	root.Right.Right = &Node{
		Val: 3,
	}
	root.Right.Right.Right = &Node{
		Val: 4,
	}
	root.Right.Right.Right.Right = &Node{
		Val: 5,
	}
	root.Right.Right.Right.Right.Right = &Node{
		Val: 6,
	}

	root.Right.Right.Child = &Node{
		Val: 7,
	}
	level2 := root.Right.Right.Child
	level2.Right = &Node{
		Val: 8,
	}
	level2.Right.Right = &Node{
		Val: 9,
	}
	level2.Right.Right.Right = &Node{
		Val: 10,
	}

	level2.Right.Child = &Node{
		Val: 11,
	}
	level2.Right.Child.Right = &Node{
		Val: 12,
	}
	r := []int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6}
	resu := flatten(root)
	for i := 0; i < len(r); i++ {
		if resu.Val != r[i] {
			t.Fail()
		}
		resu = resu.Right
	}
}
