package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	rsu := [][]int{}
    if root == nil {
        return rsu
    }
    q := []*TreeNode{root}
    var ptr *TreeNode
    var sz int
    for len(q) > 0 {
        sz = len(q)
        rsu = append([][]int{{}}, rsu...)
        for ; sz > 0; sz-- {
            ptr = q[0]
            q = q[1:]
            if ptr.Left != nil {
                q = append(q, ptr.Left)
            }
            if ptr.Right != nil {
                q = append(q, ptr.Right)
            }
            rsu[0] = append(rsu[0], ptr.Val)
        }
    }
    return rsu
}
