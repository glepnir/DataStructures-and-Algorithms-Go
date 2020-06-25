// Package main provides ...
package main

import (
	"fmt"
)

// 节点
type Node struct {
	property int
	nextNode *Node
}

// 单链表
type SingleList struct {
	len int
	// 单链表的首节点
	headNode *Node
}

func NewSingList() *SingleList {
	return &SingleList{}
}

func (s *SingleList) Display() {
	node := s.headNode
	if node == nil {
		fmt.Println("链表目前是空的")
	}
	for node != nil {
		fmt.Printf("%+v ->", node.property)
		node = node.nextNode
	}
	fmt.Println()
}

// 添加到头部方法
func (s *SingleList) AddToHead(property int) {
	node := &Node{property: property}
	if s.headNode == nil {
		s.headNode = node
	} else {
		node.nextNode = s.headNode
		s.headNode = node
	}
	s.len++
}

// 迭代遍历整个单链表
func (s *SingleList) IterateList() {
	var node *Node
	for node = s.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

// 获取最后一个节点
func (s *SingleList) LastNode() *Node {
	var node *Node
	for node = s.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			break
		}
	}
	return node
}

// 添加节点到最后
func (s *SingleList) AddToEnd(property int) {
	node := s.LastNode()
	node.nextNode = &Node{
		property: property,
	}
	s.len++
}

func main() {
	linkedList := NewSingList()
	linkedList.Display()
	fmt.Println("\n==============================")
	linkedList.AddToHead(1)
	fmt.Printf("链表的长度是%d\n", linkedList.len)
	linkedList.Display()
	fmt.Printf("当前的首节点是:%d\n", linkedList.headNode.property)
	fmt.Println("\n==============================")
	linkedList.AddToHead(3)
	fmt.Printf("链表的长度是%d\n", linkedList.len)
	linkedList.Display()
	fmt.Printf("当前的首节点是:%d\n", linkedList.headNode.property)
	fmt.Println("\n==============================")
	fmt.Println("打印整个单链表")
	linkedList.IterateList()
	fmt.Println("\n==============================")
	fmt.Println("打印最后一个节点")
	fmt.Println(linkedList.LastNode())
	fmt.Println("\n==============================")
	linkedList.AddToEnd(5)
	fmt.Printf("链表的长度是%d\n", linkedList.len)
	linkedList.Display()
	fmt.Printf("当前的首节点是:%d\n", linkedList.headNode.property)
	fmt.Println("打印整个单链表")
	linkedList.IterateList()
	fmt.Println("打印当前链表的最后一个节点")
	fmt.Println(linkedList.LastNode())
	fmt.Println("\n==============================")
}
