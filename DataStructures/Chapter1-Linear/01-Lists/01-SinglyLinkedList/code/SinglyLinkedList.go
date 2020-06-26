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

// 根据参数返回一个几点
func (s *SingleList) NodeWithValue(property int) *Node {
	node := new(Node)
	for node = s.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			break
		}
	}
	return node
}

// 在一个特殊的节点之后添加节点
func (s *SingleList) AddAfter(nodeProperty, property int) {
	node := &Node{
		property: property,
	}
	specialNode := s.NodeWithValue(nodeProperty)
	if specialNode != nil {
		node.nextNode = specialNode.nextNode
		specialNode.nextNode = node
		s.len++
	}
}

func (s *SingleList) RemoveFromEnd() {
	node := new(Node)
	if s.headNode == nil {
		fmt.Println("链表是空的")
		return
	}
	for node = s.headNode; node != nil; node = node.nextNode {
		if node.nextNode.nextNode == nil {
			node.nextNode = nil
			s.len--
		}
	}
}

func (s *SingleList) ReverseList() {
	lastnode := s.LastNode()
	s.headNode = lastnode
}

func (s *SingleList) Reverse() {
	var pre *Node
	cur := s.headNode
	if s.headNode == nil || s.headNode.nextNode == nil {
		return
	}
	for cur != nil {
		next := cur.nextNode
		cur.nextNode = pre
		pre = cur
		cur = next
	}
	s.headNode = pre
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
	fmt.Println("根据参数返回一个节点")
	fmt.Println(linkedList.NodeWithValue(1))
	fmt.Println("\n==============================")
	linkedList.AddAfter(1, 8)
	fmt.Println("添加之后的链表")
	linkedList.Display()
	fmt.Printf("当前链表的长度是%d\n", linkedList.len)
	fmt.Println("\n==============================")
	linkedList.RemoveFromEnd()
	fmt.Println("移除之后的链表")
	linkedList.Display()
	fmt.Printf("当前链表的长度是%d", linkedList.len)
	fmt.Println("\n==============================")
	linkedList.Reverse()
	linkedList.Display()
}
