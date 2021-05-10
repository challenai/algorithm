package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

// wrong answer: border case: [2->3, nil, 4, 1->6, nil]
func mergeKLists(lists []*ListNode) *ListNode {
	var ptr *ListNode
	var currentIdx int
	resu := &ListNode{}
	ptr = resu

	NotNilLeft := len(lists)

	for NotNilLeft > 0 {
		currentIdx = selectNextSmallest(lists)
		// print(lists[currentIdx].Val, " ")
		ptr.Next = lists[currentIdx]
		ptr = ptr.Next
		lists[currentIdx] = lists[currentIdx].Next
		if lists[currentIdx] == nil {
			NotNilLeft--
		}
	}

	// ptr = resu.Next
	// for ptr != nil {
	// 	print(ptr.Val, " ")
	// 	ptr = ptr.Next
	// }
	// println()

	return resu.Next
}

func selectNextSmallest(lists []*ListNode) int {
	var currentIdx int
	var currentSmall int
	currentIdx = -1
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			currentIdx = i
			currentSmall = lists[i].Val
			break
		}
	}
	for i := currentIdx + 1; i < len(lists); i++ {
		if lists[i] != nil && lists[i].Val < currentSmall {
			currentIdx = i
			currentSmall = lists[i].Val
		}
	}
	return currentIdx
}
