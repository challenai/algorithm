package solution

func rob(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var l, r int
    return robSon(root, &l, &r)
}

func robSon(root *TreeNode, l, r *int) int {
    if root == nil {
        return 0
    }
    var ll, lr, rl, rr int
    *l = robSon(root.Left, &ll, &lr)
    *r = robSon(root.Right, &rl, &rr)
    return maxInt(*l + *r, ll + lr + rl + rr + root.Val)
}

func maxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}