package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	resu := [][]int{{}}
	var ptr *TreeNode
	reverse := true
	dq := []*TreeNode{root}
	counter := 1
	for len(dq) > 0 {
		ptr = dq[0]
		dq = dq[1:]
		if reverse {
			if ptr.Right != nil {
				dq = append(dq, ptr.Right)
			}
			if ptr.Left != nil {
				dq = append(dq, ptr.Left)
			}
		} else {
			if ptr.Left != nil {
				dq = append(dq, ptr.Left)
			}
			if ptr.Right != nil {
				dq = append(dq, ptr.Right)
			}
		}
		counter--
		resu[len(resu)-1] = append(resu[len(resu)-1], ptr.Val)
		if counter == 0 {
			reverse = !reverse
			resu = append(resu, []int{})
			counter = len(dq)
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
