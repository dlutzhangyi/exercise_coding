package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
)

type Node struct {
	Key   int
	Value int
	Next  *Node
	Pre   *Node
}
type LRUCache struct {
	head    *Node
	tail    *Node
	nodeMap map[int]*Node
	cap     int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:     capacity,
		nodeMap: make(map[int]*Node),
	}
	return cache
}

func (this *LRUCache) printNodelist() {
	node := this.head
	for {
		if node == nil {
			fmt.Printf("\n")
			return
		}
		fmt.Printf("[k:%d-v:%d]->", node.Key, node.Value)
		node = node.Next
	}
}

func (this *LRUCache) moveNode2Head(node *Node) {
	if this.head == node {
		return
	}
	if this.tail == node {
		this.tail.Pre.Next = nil
		this.tail = this.tail.Pre

		node.Next = this.head
		node.Pre = nil
		this.head.Pre = node
		this.head = node
		return
	}

	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	node.Next = this.head
	node.Pre = nil
	this.head.Pre = node
	this.head = node
}

func (this *LRUCache) addNode(node *Node) {
	this.nodeMap[node.Key] = node
	if this.head == nil {
		this.head = node
		this.tail = node
		return
	}
	node.Next = this.head
	this.head.Pre = node
	this.head = node
	return
}

func (this *LRUCache) removeTailNode() {
	node := this.tail
	delete(this.nodeMap, node.Key)
	if this.tail == this.head {
		this.tail = nil
		this.head = nil
		return
	}
	this.tail.Pre.Next = nil
	this.tail = this.tail.Pre
}

//如果没有找到对应的key，则返回-1
//如果找到对应的key，则将对应的node移到head
func (this *LRUCache) Get(key int) int {
	node, ok := this.nodeMap[key]
	if !ok {
		return -1
	}
	this.moveNode2Head(node)
	return node.Value
}

//如果没有超过capacity，则将node加到head
//如果超过capacity,则移除tail，把node加到head
//如果找到这个node，则更新node的value，并移到head
func (this *LRUCache) Put(key int, value int) {
	node, ok := this.nodeMap[key]
	if ok {
		this.nodeMap[key].Value = value
		this.moveNode2Head(node)
		return
	}

	newNode := Node{
		Key:   key,
		Value: value,
	}
	if len(this.nodeMap) < this.cap {
		this.addNode(&newNode)
	} else {
		this.removeTailNode()
		this.addNode(&newNode)
	}
}

func main() {
	cache := Constructor(1)
	cache.Put(2, 1)
	v := cache.Get(2)
	if v != 1 {
		log.Errorf("v:%d", v)
	}
	cache.Put(3, 2)
	v = cache.Get(2)
	if v != -1 {
		log.Errorf("v:%d", v)
	}
	v = cache.Get(3)
	if v != 2 {
		log.Errorf("v:%d", v)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
