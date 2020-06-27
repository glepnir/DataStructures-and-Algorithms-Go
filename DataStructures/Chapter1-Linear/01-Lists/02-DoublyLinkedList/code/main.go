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

// 添加到头部方法
func (l *LinkedList) AddToHead(property int) {
	node := &Node{property: property}
	if l.headNode == nil {
		l.headNode = node
	}
	node.nextNode = l.headNode
	l.headNode.prevNode = node
	l.headNode = node
	l.len++
}

func (l *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node {
	var node *Node
	if l.len < 3 {
		fmt.Println("链表长度小于3")
		return nil
	}
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.prevNode.property == firstProperty && node.nextNode.property == secondProperty {
			break
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
}

func (l *LinkedList) AddToEnd(property int) {

}

func main() {
	l := NewLinkedList()
	l.AddToHead(7)
	l.AddToHead(4)
	fmt.Println(l.headNode.property)
	l.AddAfter(4, 9)
	fmt.Println(l.headNode.nextNode.property)
}
