package solution

type Node struct {
	Val               int
	Left, Right, Next *Node
}

func connect(root *Node) *Node {
	q := []*Node{root, nil}
	var ptr, par *Node
	for len(q) > 0 {
		if len(q) == 1 && q[0] == nil {
			break
		}
		par = ptr
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
		if par != nil {
			par.Next = ptr
		}
	}
	return root
}
