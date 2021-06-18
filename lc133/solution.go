package solution

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return node
    }
    hash := map[*Node]*Node{}
    q := []*Node{node}
    var ptr *Node
    for len(q) > 0 {
        ptr = q[0]
        q = q[1:]
        if _, ok := hash[ptr]; ok {
            continue
        }
        hash[ptr] = &Node{Val: ptr.Val, Neighbors: make([]*Node, len(ptr.Neighbors))}
        q = append(q, ptr.Neighbors...)
    }
    for n, n_ := range hash {
        for i, nList := range n.Neighbors {
            n_.Neighbors[i] = hash[nList]
        }
    }
    return hash[node]
}
