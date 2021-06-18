package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
  if root == nil {
    return [][]int{}
  }
  // 3
  // 9 20 -> 20 9
  // 1 1 15 7
  rsu := [][]int{{}}
  q := []*TreeNode{root}
  sz := 1
  reverse := true
  for len(q) > 0 {
    sz--;
    temp := q[0];
    q.unshift();
    if temp.Left != nil {
      q = append(q, temp.Left)
    }
    if temp.Right != nil {
      q = append(q, temp.Right)
    }
    if reverse {
      rsu[len(rsu)-1] = append(rsu[len(rsu)-1], temp.Val)
    } else {
      rsu[len(rsu)-1] = append([]int{temp.Val}, rsu[len(rsu)-1]...)
    }
    if sz == 0 && len(q) > 0 {
      reverse = !reverse
      rsu = append(rsu, []int{})
      sz = len(q)
    }
  }
  return rsu
}
