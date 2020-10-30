package solution

import (
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root1 := &Node{
		Val:      1,
		Children: []*Node{},
	}
	root1.Children = append(root1.Children, &Node{
		Val:      3,
		Children: []*Node{},
	})
	root1.Children = append(root1.Children, &Node{
		Val:      2,
		Children: []*Node{},
	})
	root1.Children = append(root1.Children, &Node{
		Val:      4,
		Children: []*Node{},
	})
	root1.Children[0].Children = append(root1.Children[0].Children, &Node{
		Val:      5,
		Children: []*Node{},
	})
	root1.Children[0].Children = append(root1.Children[0].Children, &Node{
		Val:      6,
		Children: []*Node{},
	})
	r1 := [][]int{{1}, {3, 2, 4}, {5, 6}}

	root2 := &Node{
		Val:      1,
		Children: []*Node{},
	}
	root2.Children = append(root2.Children, &Node{
		Val:      2,
		Children: []*Node{},
	})
	root2.Children = append(root2.Children, &Node{
		Val:      3,
		Children: []*Node{},
	})
	root2.Children = append(root2.Children, &Node{
		Val:      4,
		Children: []*Node{},
	})
	root2.Children = append(root2.Children, &Node{
		Val:      5,
		Children: []*Node{},
	})
	root2.Children[1].Children = append(root2.Children[1].Children, &Node{
		Val:      6,
		Children: []*Node{},
	})
	root2.Children[1].Children = append(root2.Children[1].Children, &Node{
		Val:      7,
		Children: []*Node{},
	})

	root2.Children[2].Children = append(root2.Children[2].Children, &Node{
		Val:      8,
		Children: []*Node{},
	})

	root2.Children[3].Children = append(root2.Children[3].Children, &Node{
		Val:      9,
		Children: []*Node{},
	})
	root2.Children[3].Children = append(root2.Children[3].Children, &Node{
		Val:      10,
		Children: []*Node{},
	})

	root2.Children[1].Children[1].Children = append(root2.Children[1].Children[1].Children, &Node{
		Val:      11,
		Children: []*Node{},
	})

	r2 := [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11}}

	rsu1 := levelOrder(root1)
	rsu2 := levelOrder(root2)
	// for i := 0; i < len(rsu1); i++ {
	// 	for j := 0; j < len(rsu1[i); j++ {
	// 		print(rsu1[i][j], " ")
	// 	}
	// 	println()
	// }
	// for i := 0; i < len(rsu2); i++ {
	// 	for j := 0; j < len(rsu2[i]); j++ {
	// 		print(rsu2[i][j], " ")
	// 	}
	// 	println()
	// }

	// t.Fail()
	if len(rsu1) != len(r1) || len(rsu2) != len(r2) {
		t.Fail()
	}
}
