## 链表

链表是用于存储项目列表的有序元素的集合，链式存储的线性表。 与数组不同，链表可以动态扩展和收缩。链表

还可以用作其他数据结构（例如堆栈和队列）的基础。 列表可用于存储用户，汽车零件，原料，待办事项和各种

其他此类元素的列表。 链表是最常用的线性数据结构。根据指针域的不同，链表分为单向链表、双向链表、循环

链表等等。

## 单链表 Linked List

单链表的特点是链表的链接方向是单向的，对链表的访问要通过从头部开始，依序往下读取。一个单链表的节点被

分为两个部分。第一个部分保存或者显示关于节点的信息，第二个部分存储下一个节点的地址，单链表只可向一个

方向遍历。不同操作在单链表的时间复杂度

```
Access(存取) ---- O(n)
Search(搜索) ---- O(n)
Insert(插入) ---- O(1)
Delete(删除) ---- O(1)
```

### Node 节点

Node 节点是一个结构体 Struct，结构体的第一个成员是个整形 int 的字段 property，这里根据你的需求来不一

定是整形也不一定要叫 property:)。第二个成员是指向下一个节点的指针。

```Go
type Node struct{
  property int
  nextNode *node
}
```

### 定义单链表

```GO
type LinkedList struct{
  // 链表的changdu
  len int
  // 单链表的首节点
  headNode *Node
}
```

下面列出 LinkedList 的一些方法，例如 AddtoHead，IterateList，LastNode，AddtoEnd，NodeWithValue，AddAfter

### AddToHead 方法

> 写代码很重要的是要先想好思路。所以当你在看本教程的时候，请先理解思路在看代码示例。

AddToHead 方法将节点添加到单链表的开头。理一下思路每一步要做什么代码也就自然的写出来了。

- 首先我们的单链表里有首节点的信息字段 headNode。第一步要判断当前的首节点是不是 nil，如果是 nil 那么
  当前的链表是空的，那么直接将节点赋值到 headNode。
- 当首节点不为空的话，第一步要将首节点的下一个节点信息 headNode.nextNode 赋值给新节点的 nextNode，将这
  个新节点赋值给 headNode。

示例代码:

```go
// 节点
type Node struct {
  // 节点的属性
	property int
  // 指向下一个节点的指针
	nextNode *Node
}

// 单链表
type SingleList struct {
  //链表的长度
	len int
	// 单链表的首节点
	headNode *Node
}

// 简单的工厂函数用于返回一个初始化的
// 单链表指针
func NewSingList() *SingleList {
	return &SingleList{}
}

// 用于显示整个链表
func (s *SingleList) Display() {
	node := s.headNode
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
func (s *SingleList) AddToHead(property int) {
  // 根据传入的属性值生成一个新节点
  // 这个新节点的property就是这个方法的参数
  // 这个新节点的nextNode是nil
	node := &Node{property: property}
  // 判断当前链表的首节点是nil代表链表是空的
	if s.headNode == nil {
    // 空链表直接赋值到headNode
		s.headNode = node
	} else {
    // 不是空链表将当前首节点里保存的指向下一个节点的指针
    // 赋值给新节点的nextNode指针，让这个新节点的指向下一
    // 个节点的指针指向当前首节点指向的下一个节点
		node.nextNode = s.headNode
    // 然后将这个新节点赋值给首节点 完成了在头部插入节点
		s.headNode = node
	}
  // 链表长度加1
	s.len++
}
```

- 在 main 函数中运行一下

```go

func main() {
	linkedList := NewSingList()
	linkedList.Display()
	fmt.Println("\n==============================\n")
	linkedList.AddToHead(1)
	fmt.Printf("链表的长度是%d\n", linkedList.len)
	linkedList.Display()
	fmt.Printf("当前的首节点是:%d\n", linkedList.headNode.property)
	fmt.Println("\n==============================\n")
	linkedList.AddToHead(3)
	fmt.Printf("链表的长度是%d\n", linkedList.len)
	linkedList.Display()
	fmt.Printf("当前的首节点是:%d", linkedList.headNode.property)
}
```

- 输出
  ![](/image/linkedlist/01.png)

## IterateList 方法迭代遍历整个单链表

- 遍历节点的思路其实和 for 循环的思路差不多。一般我们简单的 for 循环是这样 for 里定义个 i 然后给 i 加
  条件不满足就增加 i 的值，迭代遍历单链表其实也差不多只不过这个 i 我们定义一个新节点然后把首节点赋值
  给它，结束的条件？什么时候跳出 for 循环？很明显到单链表最后一个节点，最后一个节点里的 nextNode 是
  nil 的。所以不满足时这个 node 的变化就是依次的每个节点的 nextNode 赋值给新定义的 node。直到这个 node
  的值是 nil 也就是到了单链表的最后了。

```go
func (s *SingleList) IterateList() {
  // 定义一个新节点
	var node *Node
  // 将单链表的首节点赋值给它，当这个node不为nil是把nextNode的值赋给node
  // 知道这个node为nil
	for node = s.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}
```
