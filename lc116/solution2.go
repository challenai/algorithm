package solution

type Node struct {
	Val               int
	Left, Right, Next *Node
}

func connect1(root *Node) *Node {
    if root == nil {
        return root
    }
    connect2(root.Left, root.Right)
    return root
}

func connect2(l, r *Node) {
    if l == nil || r == nil {
        return
    }
    l.Next = r
    connect2(l.Left, l.Right)
    connect2(l.Right, r.Left)
    connect2(r.Left, r.Right)
}
