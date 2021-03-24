package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func FindBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var resu int
	var current *TreeNode
	q := []*TreeNode{root, nil}
	for len(q) > 1 && q[0] != nil {
		current = q[0]
		q = q[1:]
		if current.Left != nil {
			q = append(q, current.Left)
		}
		if current.Right != nil {
			q = append(q, current.Right)
		}
		if q[0] == nil {
			q = q[1:]
			if len(q) > 0 && q[0] != nil {
				resu = q[0].Val
			}
			q = append(q, nil)
		}
	}
	return resu
}
