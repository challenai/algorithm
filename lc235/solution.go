package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// idea: store all the path and put it in a hash
	if q == p || p == nil || q == nil {
		return root
	}
	hash := map[*TreeNode]bool{}
	var ptr, lastCorresponseNode *TreeNode
	ptr = root
	for ptr != nil {
		if ptr.Val == p.Val {
			break
		}
		hash[ptr] = false
		if p.Val > ptr.Val {
			ptr = ptr.Right
		} else {
			ptr = ptr.Left
		}
	}

	ptr = root
	lastCorresponseNode = nil
	for ptr != nil {
		if _, ok := hash[ptr]; ok {
			lastCorresponseNode = ptr
		} else {
			return lastCorresponseNode
		}
		if q.Val > ptr.Val {
			ptr = ptr.Right
		} else {
			ptr = ptr.Left
		}
	}

	return lastCorresponseNode
}
