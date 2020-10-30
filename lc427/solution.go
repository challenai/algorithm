package solution

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var current []int
	var ptr *Node
	resu := [][]int{}
	if root == nil {
		return resu
	}
	q := []*Node{root, nil}
	for len(q) > 0 {
		ptr = q[0]
		q = q[1:]
		if ptr == nil {
			temp := make([]int, len(current))
			copy(temp, current)
			resu = append(resu, temp)
			current = []int{}
			if len(q) > 0 {
				q = append(q, nil)
			}
			continue
		}
		current = append(current, ptr.Val)
		if ptr.Children != nil {
			q = append(q, ptr.Children...)
		}
	}
	return resu
}
