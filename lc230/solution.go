package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
	// for frequent query and modify
	// ChildrenCnt int
}

func kthSmallest(root *TreeNode, k int) int {
	ptr := root
	st := []*TreeNode{}
	counter := 0
	for ptr != nil || len(st) > 0 {
		if ptr != nil {
			st = append(st, ptr)
			ptr = ptr.Left
		} else {
			ptr = st[len(st)-1]
			st = st[:len(st)-1]
			counter++
			if counter == k {
				return ptr.Val
			}
			ptr = ptr.Right
		}
	}
	return 0
}
