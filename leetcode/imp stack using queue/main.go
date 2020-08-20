package main

import "fmt"

type Queue struct {
	queue []int
}

func NewQueue() *Queue {
	q := Queue{}
	q.queue = []int{}
	return &q
}

func (q *Queue) Push(v int) {
	q.queue = append(q.queue, v)
}

func (q *Queue) Pop() int {
	if q.Empty() {
		return 0
	}
	v := q.queue[0]
	if len(q.queue) == 1 {
		q.queue = q.queue[:0]
	} else {
		q.queue = q.queue[1:]
	}
	return v
}

func (q *Queue) Empty() bool {
	if len(q.queue) == 0 {
		return true
	}
	return false
}

type MyStack struct {
	q1, q2 *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	s := MyStack{}
	s.q1, s.q2 = NewQueue(), NewQueue()
	return s
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.q1.Empty() {
		this.q2.Push(x)
	} else {
		this.q1.Push(x)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	v := 0
	if this.q1.Empty() {
		for !this.q2.Empty() {
			v = this.q2.Pop()
			if this.q2.Empty() {
				return v
			} else {
				this.q1.Push(v)
			}
		}
	} else {
		for !this.q1.Empty() {
			v = this.q1.Pop()
			if this.q1.Empty() {
				return v
			} else {
				this.q2.Push(v)
			}
		}
	}
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	v := 0
	if this.q1.Empty() {
		for !this.q2.Empty() {
			v = this.q2.Pop()
			this.q1.Push(v)
		}
	} else {
		for !this.q1.Empty() {
			v = this.q1.Pop()
			this.q2.Push(v)
		}
	}
	return v
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.q1.Empty() && this.q2.Empty() {
		return true
	}
	return false
}

func main() {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
}
