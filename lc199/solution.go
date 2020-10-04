package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	resu := []int{}
	q := []*TreeNode{root, nil}
	var ptr *TreeNode
	for len(q) > 1 {
		ptr = q[0]
		q = q[1:]
		if ptr == nil {
			continue
		}
		if ptr.Left != nil {
			q = append(q, ptr.Left)
		}
		if ptr.Right != nil {
			q = append(q, ptr.Right)
		}
		if q[0] == nil {
			resu = append(resu, ptr.Val)
			q = q[1:]
			q = append(q, nil)
		}
	}
	return resu
}
