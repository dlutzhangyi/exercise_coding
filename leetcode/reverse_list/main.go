package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var before, current, next *ListNode
	before = nil
	current = head
	next = head.Next

	for current != nil {
		fmt.Println(current.Val)
		current.Next = before
		before = current
		current = next
		if current == nil {
			return before
		}
		next = next.Next
	}
	return before
}

func buildList(values []int) *ListNode {
	var head, node *ListNode
	head = &ListNode{
		Val: values[0],
	}
	node = head
	for i := 1; i < len(values); i++ {
		tmp := &ListNode{
			Val: values[i],
		}
		node.Next = tmp
		node = node.Next
	}
	return head
}

func printNodelist(head *ListNode) {
	node := head
	for node != nil {
		fmt.Printf("%d->", node.Val)
		node = node.Next
	}
	fmt.Println("\n")
}
func main() {
	head := buildList([]int{1, 2, 3, 4})
	printNodelist(head)
	reverseList(head)
}
