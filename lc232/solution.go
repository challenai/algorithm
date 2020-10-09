package solution

import (
	"lc2/utils"
)

type MyQueue struct {
	sts [2]*utils.Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		sts: [2]*utils.Stack{
			utils.NewStack(),
			utils.NewStack(),
		},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.sts[0].Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.sts[0].IsEmpty() {
		return 0
	}
	var temp int
	for !this.sts[0].IsEmpty() {
		temp = this.sts[0].Pop()
		if !this.sts[0].IsEmpty() {
			this.sts[1].Push(temp)
		}
	}
	for !this.sts[1].IsEmpty() {
		this.sts[0].Push(this.sts[1].Pop())
	}
	return temp
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.sts[0].IsEmpty() {
		return 0
	}
	var temp int
	for !this.sts[0].IsEmpty() {
		temp = this.sts[0].Pop()
		this.sts[1].Push(temp)
	}
	for !this.sts[1].IsEmpty() {
		this.sts[0].Push(this.sts[1].Pop())
	}
	return temp
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.sts[0].IsEmpty()
}
