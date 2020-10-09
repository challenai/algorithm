package utils

type node struct {
	val  int
	next *node
}

type Stack struct {
	head *node
}

func NewStack() *Stack {
	return &Stack{}
}

func (st *Stack) Push(v int) {
	st.head = &node{
		val:  v,
		next: st.head,
	}
}

func (st *Stack) Pop() int {
	if st.head == nil {
		return 0
	}
	temp := st.head.val
	st.head = st.head.next
	return temp
}

func (st *Stack) Peek() int {
	if st.head == nil {
		return 0
	}
	return st.head.val
}

func (st *Stack) IsEmpty() bool {
	return st.head == nil
}
