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
