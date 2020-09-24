package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	resu := [][]int{}
	st := []*TreeNode{root, nil}
	q := []*TreeNode{root, nil}
	var ptr *TreeNode

	for len(q) > 0 {
		ptr = q[0]
		q = q[1:]
		if ptr == nil {
			if len(q) == 0 {
				break
			}
			q = append(q, nil)
			st = append(st, nil)
			continue
		}
		if ptr.Right != nil {
			q = append(q, ptr.Right)
			st = append(st, ptr.Right)
		}
		if ptr.Left != nil {
			q = append(q, ptr.Left)
			st = append(st, ptr.Left)
		}
	}
	for len(st) > 0 {
		ptr = st[len(st)-1]
		st = st[:len(st)-1]
		if ptr == nil {
			resu = append(resu, []int{})
			continue
		}
		resu[len(resu)-1] = append(resu[len(resu)-1], ptr.Val)
	}
	// no test so just redirect to monitor to check out
	// for i := 0; i < len(resu); i++ {
	// 	for j := 0; j < len(resu[i]); j++ {
	// 		print(resu[i][j], " ")
	// 	}
	// 	println()
	// }

	return resu
}
