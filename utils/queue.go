package utils

type Queue struct {
	head *node
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
	}
}

func (q *Queue) Enque(v int) {
	if q.head == nil {
		q.head = &node{
			val: v,
		}
		return
	}
	ptr := q.head
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = &node{
		val: v,
	}
}

func (q *Queue) Deque() int {
	if q.head == nil {
		return 0
	}
	temp := q.head.val
	q.head = q.head.next
	return temp
}

func (q *Queue) Peek() int {
	if q.head == nil {
		return 0
	}
	return q.head.val
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}
