package solution

import "lc2/utils"

type MyStack struct {
	qs [2]*utils.Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		qs: [2]*utils.Queue{
			utils.NewQueue(),
			utils.NewQueue(),
		},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.qs[0].Enque(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.qs[0].IsEmpty() {
		return 0
	}
	var temp int
	for !this.qs[0].IsEmpty() {
		temp = this.qs[0].Deque()
		if !this.qs[0].IsEmpty() {
			this.qs[1].Enque(temp)
		}
	}
	this.qs[0], this.qs[1] = this.qs[1], this.qs[0]
	return temp
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.qs[0].IsEmpty() {
		return 0
	}
	var temp int
	for !this.qs[0].IsEmpty() {
		temp = this.qs[0].Deque()
		this.qs[1].Enque(temp)
	}
	this.qs[0], this.qs[1] = this.qs[1], this.qs[0]
	return temp
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.qs[0].IsEmpty()
}
