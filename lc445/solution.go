package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(list1, list2 *ListNode) *ListNode {
	var rsu *ListNode
	rsu = nil
	st1 := []*ListNode{}
	st2 := []*ListNode{}
	ptr1 := list1
	ptr2 := list2
	for ptr1 != nil {
		st1 = append(st1, ptr1)
		ptr1 = ptr1.Next
	}
	for ptr2 != nil {
		st2 = append(st2, ptr2)
		ptr2 = ptr2.Next
	}
	var current int
	needPlusOne := 0
	for len(st1) > 0 || len(st2) > 0 {
		if len(st1) > 0 && len(st2) > 0 {
			println(st1[len(st1)-1].Val, st2[len(st2)-1].Val)
			current = st1[len(st1)-1].Val + st2[len(st2)-1].Val + needPlusOne
			st1 = st1[:len(st1)-1]
			st2 = st2[:len(st2)-1]
		} else if len(st1) > 0 {
			current = st1[len(st1)-1].Val + needPlusOne
			st1 = st1[:len(st1)-1]
		} else {
			current = st2[len(st2)-1].Val + needPlusOne
			st2 = st2[:len(st2)-1]
		}
		// bug here, lack reset needPlusOne
		needPlusOne = 0
		if current > 9 {
			current -= 10
			needPlusOne = 1
		}
		rsu = &ListNode{Val: current, Next: rsu}
	}
	if needPlusOne > 0 {
		rsu = &ListNode{Val: 1, Next: rsu}
	}
	return rsu
}
