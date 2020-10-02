package solution

type node struct {
	val, min int
	next     *node
}

type MinStack struct {
	head *node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		head: nil,
	}
}

func (this *MinStack) Push(x int) {
	temp := &node{
		val:  x,
		min:  x,
		next: this.head,
	}
	this.head = temp
	if this.head.next != nil && this.head.next.min < this.head.min {
		this.head.min = this.head.next.min
	}
}

func (this *MinStack) Pop() {
	if this.head == nil {
		return
	}
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	if this.head == nil {
		return 0
	}
	return this.head.val
}

func (this *MinStack) GetMin() int {
	if this.head == nil {
		return 0
	}
	return this.head.min
}
