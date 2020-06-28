// Package main provides ...
package main

import "fmt"

// 双链表节点
// 相比单链表多了一个prevNode指向前一个节点
type Node struct {
	property int
	nextNode *Node
	prevNode *Node
}

type LinkedList struct {
	len      int
	headNode *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// 用于显示整个链表
func (l *LinkedList) Display() {
	node := l.headNode
	if node == nil {
		fmt.Println("链表目前是空的")
	}
	// 循环遍历节点直到节点是nil为止 代表到头了
	for node != nil {
		fmt.Printf("%+v ->", node.property)
		node = node.nextNode
	}
	fmt.Println()
}

// 添加到头部方法
func (l *LinkedList) AddToHead(property int) {
	node := &Node{property: property}
	if l.headNode == nil {
		l.headNode = node
		l.len++
	} else {
		node.nextNode = l.headNode
		l.headNode.prevNode = node
		l.headNode = node
		l.len++
	}
}

func (l *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node {
	var node *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.prevNode != nil && node.nextNode != nil {
			if node.prevNode.property == firstProperty && node.nextNode.property == secondProperty {
				break
			}
		}
	}
	return node
}

func (l *LinkedList) NodeWithValue(property int) *Node {
	node := new(Node)
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			break
		}
	}
	return node
}

func (l *LinkedList) AddAfter(nodeProperty, property int) {
	node := &Node{property: property}
	specialNode := l.NodeWithValue(nodeProperty)
	node.prevNode = specialNode
	node.nextNode = specialNode.nextNode
	specialNode.nextNode = node
	l.len++
}

func (l *LinkedList) LastNode() *Node {
	node := new(Node)
	switch l.len {
	case 0:
		fmt.Println("linked list is empty")
		return nil
	case 1:
		return l.headNode
	}
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			break
		}
	}
	return node
}

func (l *LinkedList) AddToEnd(property int) {
	node := &Node{property: property}
	lastNode := l.LastNode()
	lastNode.nextNode = node
	node.prevNode = lastNode
}

func main() {
	l := NewLinkedList()
	l.AddToHead(7)
	l.AddToHead(4)
	fmt.Println(l.headNode.property)
	l.AddAfter(4, 9)
	l.Display()
	fmt.Println(l.NodeBetweenValues(4, 7))
	lastnode := l.LastNode()
	fmt.Println(lastnode)
	l.AddToEnd(6)
	l.Display()
}
