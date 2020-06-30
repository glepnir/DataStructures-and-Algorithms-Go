## CircularLinkedList 环形链表

环形链表可以是单链表也可以是双链表，单链表的最后一个节点的 nextNode 是指向 nil 的，如果将它指向首节点
那么久形成了一个简单的环形链表，对于双向链表将首节点的 prevNode 指向最后一个节点，将最后一个节点的
nextNode 指向首节点，也就形成了环形的双向链表。

## 环形单链表

首先定义一个环形单链表,这里和第一节的单链表定义并无区别。注意的是当环形链表只有一个节点的时候，它指
向的是自己。

```go

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
```

### AddToHead 添加到头部的方法

- 大概的思路是和单链表类似的，区别在于添加之后我们需要找到最后一个节点将最后一个节点的 nextNode 指针
- 指向新插入的头结点。

```GO
func (c *CircularSingleList) AddToHead(property int) {
	node := &Node{property: property}
	if c.len == 0 {
		c.headNode = node
		c.headNode.nextNode = c.headNode
		c.len++
	} else {
		lastNode := new(Node)
    // 找到最后一个节点
		for lastNode = c.headNode; ; lastNode = lastNode.nextNode {
			if lastNode.nextNode == c.headNode {
				break
			}
		}
    // 新节点的下一个节点指向当前的头结点
		node.nextNode = c.headNode
    // 最后一个节点指向新的节点
		lastNode.nextNode = node
    // 将新节点赋值到首节点完成插入到头部
		c.headNode = node
		c.len++
	}
}

```

- 输出

```go
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
}
```

### LastNode 查找最后一个节点的方法

- 直接从 AddToHead 中抽取出来即可。

```GO
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
```

### AddToEnd 尾部增加的方法

- 通过 LastNode 方法获取到最后一个节点，因为这个新节点是要放到最后的，所以这个新节点的 nextNode
- 一定是指向头结点的直接赋值就可以了。然后将当前的 lastNode 的 nextNode 指向这个新节点

```go
func (c *CircularSingleList) AddToEnd(property int) {
	node := &Node{property: property}
	lastNode := c.LastNode()
	node.nextNode = c.headNode
	lastNode.nextNode = node
}
```

### RemoveFromEnd 移除尾部节点

- 移除尾部节点需要找到最后一个节点的前一个节点，使其指向头结点就移除了最后的节点

```go
func (c *CircularSingleList) RemoveFromEnd() {
	node := new(Node)
	for node = c.headNode; ; node = node.nextNode {
		if node.nextNode.nextNode == c.headNode {
			break
		}
	}
	node.nextNode = c.headNode
}

```

> 最主要是在脑海里要有链表的形状理清楚逻辑那么例如删除特定节点或在特定节点插入都是很简单的。

### 输出

```GO

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
	c.RemoveFromEnd()
	fmt.Println(c.LastNode().property) // 输出7
}
```

## 环形双向链表

环形双向链表是在双链表的基础上，链表的最后一个节点指向首节点，首节点的 prevNode 指向链表的最后一个节点

```go
type Node struct {
	property int
	nextNode *Node
	prevNode *Node
}

type CircularDoubleList struct {
	len      int
	headNode *Node
}

func NewCircularDoubleList() *CircularDoubleList {
	return &CircularDoubleList{}
}
```

### AddToHead 添加节点到头部

- 相比环形单链表，环形双向链表的 AddToHead 方法的区别在于需要处理新增节点的 prevNode 的处理，它应该
- 指向当前的尾部节点

```go

func (c *CircularDoubleList) AddToHead(property int) {
	node := &Node{property: property}
	if c.len == 0 {
		node.nextNode = node
		node.prevNode = node
		c.headNode = node
		c.len++
	} else {
		lastNode := new(Node)
		for lastNode = c.headNode; ; lastNode = node.nextNode {
			if lastNode.nextNode == c.headNode {
				break
			}
		}
		lastNode.nextNode = node
		node.nextNode = c.headNode
    // 将新的节点的prevNode指向当前的尾部节点
		node.prevNode = lastNode
		c.headNode = node
		c.len++
	}
}

```

- 输出

```GO
func main() {
	c := NewCircularDoubleList()
	c.AddToHead(9)
	// 打印当前的首节点正确输出是9
	fmt.Println(c.headNode.property) // 输出9
	c.AddToHead(5)
	// 打印当前的首节点正确输出是5
	fmt.Println(c.headNode.property) // 输出5
	// 因为当前只有2个节点所以首节点的nextNode也就是尾部节点
	// 正确输出是9
	fmt.Println(c.headNode.nextNode.property) // 输出9
	// 当前最后节点的下一个节点指向应该是首节点环形
	// 正确输出是首节点的5
	fmt.Println(c.headNode.nextNode.nextNode.property) // 输出 5
	// 首节点的前一个节点指向应该是尾结点
	// 所以正确输出应该是9
	fmt.Println(c.headNode.prevNode.property) // 输出 9
}
```

项目地址:[github:DataStructuresAndAlgorithms-Go](https://github.com/glepnir/DataStructuresAndAlgorithms-Go)
如果喜欢这个项目请 Star。
