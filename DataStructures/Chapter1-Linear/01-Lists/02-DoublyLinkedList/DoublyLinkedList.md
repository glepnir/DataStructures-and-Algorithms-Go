## 双链表

双链表是在单链表的基础上再 node 节点添加一个指向前一个节点的指针。在双向链表中，所有节点在列表的任一
侧都有指向它们所连接的节点的指针。 这意味着每个节点都连接到两个节点，我们可以向前遍历到下一个节点，
也可以向后遍历到上一个节点。 双链表允许插入，删除和遍历等操作

![](/image/linkedlist/05.png)

```GO
type Node struct{
  property int
  nextNode *Node
  prevNode *Node
}

// 双链表
type LinkedList struct {
	len      int
	headNode *Node
}
```

## AddToHead

- 添加到双链表的头部，这个方法与单链表的有一点点差异
- 在双链表中每个节点有一个 prevNode 指向前一个节点，所
- 以要插入到头部那就是让这个原始的头结点 prevNode 指向这个
- 新节点，新节点的 nextNode 指向这个头结点就完成了头部插入

```go
// 添加到头部方法
func (l *LinkedList) AddToHead(property int) {
	node := &Node{property: property}
  // 判断是否为空
	if l.headNode == nil {
		l.headNode = node
	}
  // 新的头结点的nextNode指向下一个节点的指针指向现在双链表的
  // 的头结点，现在双链表的prevNode指向这个新节点。
	node.nextNode = l.headNode
	l.headNode.prevNode = node
  // 将新节点赋值给头结点
	l.headNode = node
  // 双链表长度加1
	l.len++
}
```

## NodeBetweenValues

- 这个方法是通过传入的 2 个参数找到位于它们中间的节点。理解了单链表的 NodeWithValue 方法
- 这个方法也就理解了是单链表的 NodeWithValue 的变种只是多加了一个前节点的判断即可。

```go
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
```

##
