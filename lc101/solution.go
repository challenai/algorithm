package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric2(root.Left, root.Right)
}

func isSymmetric2(p, q *TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val && isSymmetric2(p.Left, q.Right) && isSymmetric2(p.Right, q.Left)
	}
	if p == nil && q == nil {
		return true
	}
	return false
}
