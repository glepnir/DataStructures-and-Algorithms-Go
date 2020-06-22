// Package main provides ...
package main

import "fmt"

// 节点
type Node struct {
	property int
	nextNode *Node
}

// 单链表
type LinkedList struct {
	// 单链表的首节点
	headNode *Node
}

// 添加到头部方法
func (linkedList *LinkedList) AddToHead(property int) {
	node := Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

func main() {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	fmt.Println(linkedList.headNode.property)
}
