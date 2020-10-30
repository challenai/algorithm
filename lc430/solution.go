package solution

type Node struct {
	Val                int
	Left, Right, Child *Node
}

func flatten(root *Node) *Node {
	ptr := root
	st := []*Node{}
	for ptr != nil || len(st) > 0 {
		// println(ptr.Val)
		if ptr.Child != nil {
			for ptr.Child != nil {
				st = append(st, ptr.Right)
				ptr.Right = ptr.Child
				ptr.Right.Left = ptr
				ptr.Child = nil
			}
		}
		if ptr.Right == nil && len(st) > 0 {
			ptr.Right = st[len(st)-1]
			st = st[:len(st)-1]
			ptr.Right.Left = ptr
		}
		ptr = ptr.Right
	}
	return root
}
