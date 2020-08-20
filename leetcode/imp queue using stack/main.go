package main

import (
	"fmt"
)

type MyQueue struct {
	stack1, stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	q := MyQueue{}
	q.stack1, q.stack2 = []int{}, []int{}
	return q
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) move() {
	for i := len(this.stack1) - 1; i >= 0; i-- {
		this.stack2 = append(this.stack2, this.stack1[i])
	}
	this.stack1 = this.stack1[:0]
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return 0
	}
	value := 0
	//stack2为空，将stack1中的数据导入到stack2中
	if len(this.stack2) == 0 {
		this.move()
	}
	value = this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[0 : len(this.stack2)-1]
	return value
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return 0
	}
	//stack2为空，将stack1中的数据导入到stack2中
	if len(this.stack2) == 0 {
		this.move()
	}
	return this.stack2[len(this.stack2)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return true
	}
	return false
}

func main() {

	queue := Constructor()

	queue.Push(1)
	fmt.Println(queue.Pop())   // returns 1
	fmt.Println(queue.Empty()) // returns false
}
