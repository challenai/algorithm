package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func LargestValue(root *TreeNode) []int {
	resu := []int{}
	if root == nil {
		return resu
	}

	resu = append(resu, root.Val)
	var current *TreeNode
	q := []*TreeNode{root, nil}
	for len(q) > 1 {
		current = q[0]
		q = q[1:]
		if current.Val > resu[len(resu)-1] {
			resu[len(resu)-1] = current.Val
		}
		if current.Left != nil {
			q = append(q, current.Left)
		}
		if current.Right != nil {
			q = append(q, current.Right)
		}
		if q[0] == nil {
			q = q[1:]
			if len(q) > 0 && q[0] != nil {
				resu = append(resu, q[0].Val)
			}
			q = append(q, nil)
		}
	}
	return resu
}
