package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func genNode(list []int) *ListNode {
	if len(list) == 0 {
		return &ListNode{}
	}
	head := &ListNode{
		Val: list[0],
	}
	var p, q *ListNode
	p = head

	for i := 1; i < len(list); i++ {
		q = &ListNode{
			Val: list[i],
		}
		p.Next = q
		p = q
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, node *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	node = head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}

	return head

}
func printlist(list *ListNode) {
	if list == nil {
		return
	}
	for {
		if list == nil {
			break
		}
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}

}
func main() {
	list1 := genNode([]int{2, 6})
	list2 := genNode([]int{1, 3, 4})
	list := mergeTwoLists(list1, list2)
	printlist(list)
}
