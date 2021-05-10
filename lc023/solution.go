package solution

// haha only beats 41.85%~~~
func mergeKLists(lists []*ListNode) *ListNode {
  rsu := &ListNode{
    Next: nil,
  }
  ptr := rsu
  var current *ListNode
  var currentIdx int
  for i := 0; i < len(lists); i++ {
    for i < len(lists) && lists[i] == nil {
      lists[i], lists[len(lists)-1] = lists[len(lists)-1], lists[i]
      lists = lists[:len(lists)-1]
    }
  }
  for len(lists) > 0 {
    // select the smallest one
    current = lists[0]
    currentIdx = 0
    for i := 1; i < len(lists); i++ {
      if lists[i].Val < current.Val {
        current = lists[i]
        currentIdx = i
      }
    }
    lists[currentIdx] = current.Next
    ptr.Next = current
    ptr = ptr.Next
    if current.Next == nil {
      lists[currentIdx], lists[len(lists)-1] = lists[len(lists)-1], lists[currentIdx]
      lists = lists[:len(lists)-1]
    }
  }
  return rsu.Next
}
