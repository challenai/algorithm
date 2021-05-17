package solution

func generateTrees(n int) []*TreeNode {
  if n == 0 {
    return []*TreeNode{}
  }
  return loop(1, n)
}

func loop(lo, hi int) []*TreeNode {
  if lo > hi {
    return []*TreeNode{nil}
  }
  rsu := []*TreeNode{}
  for i := lo; i <= hi; i++ {
    left := loop(lo, i-1)
    right := loop(i+1, hi)
    for _, t1 := range left {
      for _, t2 := range right {
        root := &TreeNode{
          Val: i,
        }
        root.Left = t1
        root.Right = t2
        rsu = append(rsu, root)
      }
    }
  }
  return rsu
}
