package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	resu := [][]int{{}}
	q := []*TreeNode{}
	counter := 1
	var ptr *TreeNode
	q = append(q, root)
	for len(q) > 0 {
		ptr = q[0]
		q = q[1:]
		if ptr.Right != nil {
			q = append(q, ptr.Right)
		}
		if ptr.Left != nil {
			q = append(q, ptr.Left)
		}
		resu[len(resu)-1] = append(resu[len(resu)-1], ptr.Val)
		counter--
		if counter == 0 {
			counter = len(q)
			resu = append(resu, []int{})
		}
	}
	// for i := 0; i < len(resu); i++ {
	// 	for j := 0; j < len(resu[i]); j++ {
	// 		print(resu[i][j], " ")
	// 	}
	// 	println()
	// }

	return resu
}
