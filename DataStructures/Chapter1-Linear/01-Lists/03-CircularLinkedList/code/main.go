// Package main provides ...
package main

import "fmt"

// 节点
type Node struct {
	property int
	nextNode *Node
}

type CircularSingleList struct {
	len int
	// 环形单链表的首节点
	headNode *Node
}

func NewCircularSingleList() *CircularSingleList {
	return &CircularSingleList{}
}

func (c *CircularSingleList) AddToHead(property int) {
	node := &Node{property: property}
	if c.len == 0 {
		c.headNode = node
		c.headNode.nextNode = c.headNode
		c.len++
	} else {
		lastNode := new(Node)
		for lastNode = c.headNode; ; lastNode = lastNode.nextNode {
			if lastNode.nextNode == c.headNode {
				break
			}
		}
		node.nextNode = c.headNode
		lastNode.nextNode = node
		c.headNode = node
		c.len++
	}
}

func (c *CircularSingleList) LastNode() *Node {
	node := new(Node)
	if c.len == 0 {
		return nil
	}
	for node = c.headNode; ; node = node.nextNode {
		if node.nextNode == c.headNode {
			break
		}
	}
	return node
}

func (c *CircularSingleList) AddToEnd(property int) {
	node := &Node{property: property}
	lastNode := c.LastNode()
	node.nextNode = c.headNode
	lastNode.nextNode = node
}

func main() {
	// 初始化一个空的环形单链表
	c := NewCircularSingleList()
	// 添加到头部 此时的环形单链表首节点是7
	c.AddToHead(7)
	// 再添加一个节点到头部 那么此时的首节点是8
	c.AddToHead(8)
	// 打印当前的首节点 输出8
	fmt.Println(c.headNode.property)
	// 打印当前的环形单链表的第二个节点也是当前环形链表的最后一个节点应该是7
	fmt.Println(c.headNode.nextNode.property) // 输出7
	// 打印最后一个节点的下一个节点的指向 应该是指向首节点也就是8
	fmt.Println(c.headNode.nextNode.nextNode.property) // 输出8
	c.AddToEnd(4)
	fmt.Println(c.LastNode().property) // 输出4
}
