## 链表

链表是用于存储项目列表的有序元素的集合，链式存储的线性表。 与数组不同，链表可以动态扩展和收缩。链表

还可以用作其他数据结构（例如堆栈和队列）的基础。 列表可用于存储用户，汽车零件，原料，待办事项和各种

其他此类元素的列表。 链表是最常用的线性数据结构。根据指针域的不同，链表分为单向链表、双向链表、循环

链表等等。

## 单链表 Linked List

单链表的特点是链表的链接方向是单向的，对链表的访问要通过从头部开始，依序往下读取。一个单链表的节点被

分为两个部分。第一个部分保存或者显示关于节点的信息，第二个部分存储下一个节点的地址，单链表只可向一个

方向遍历。不同操作在单链表的时间复杂度

|  动 作   | 数 组 | 链 表 |
| :------: | :---: | :---: |
|  存 取   | O(1)  | O(n)  |
|  搜 索   | O(1)  | O(n)  |
| 头 增 加 | O(1)  | O(1)  |
| 尾 增 加 | O(n)  | O(1)  |
|  删 除   | O(n)  | O(n)  |

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

下面列出 LinkedList 的一些方法实现。

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

### IterateList 方法迭代遍历整个单链表

> 写代码很重要的是要先想好思路。所以当你在看本教程的时候，请先理解思路在看代码示例。

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
  // 直到这个node为nil
	for node = s.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}
```

- 输出
  ![](/image/linkedlist/02.png)

### LastNode 方法获得链表的最后一个元素

> 写代码很重要的是要先想好思路。所以当你在看本教程的时候，请先理解思路在看代码示例。

- 获得最后一个节点思路只要你理解了上一个遍历也就很容易理解了。当 node 为 nil 就会跳出 for 循环，那就
  很简单了判断节点的下一个节点 nextNode 是 nil 不就找到最后一个节点了吗，所以代码就是上个例子的变种了

```go
func (s *SingleList) LastNode() *Node {
	var node *Node
  // 将首节点赋值给新定义的节点，当节点为nil的时候跳出 for 循环
  // node的条件每次将当前节点的nextNode赋值给它
	for node = s.headNode; node != nil; node = node.nextNode {
    // 如果当前这个node的nextNode为nil就是没有指向下一个节点的指针
    // 这个节点就是最后一个节点跳出 for 循环
		if node.nextNode == nil {
			break
		}
	}
  // 返回当前的这个node
	return node
}
```

- 输出就不贴图了。。截图有点麻烦。。末尾贴出所有的代码和运行截图

### AddToEnd 添加到最后的方法

- 这个更简单了。上面都写出来了取得当前单链表最后一个节点的方法。直接将 LastNode 方法取到的最后一个节点
  的 nextNode 的值设置成新创建的节点不就添加到最后了吗。代码也就写出来了。

```Go
func (s *SingleList) AddToEnd(property int) {
  // 通过LastNode方法获取到最后一个节点
	node := s.LastNode()
  // 将当前最后一个节点的nextNode赋值
	node.nextNode = &Node{
		property: property,
	}
  // 链表长度加1
	s.len++
}
```

### NodeWithValue 根据参数返回一个节点

- 根据传入的参数在单链表中找到这个节点并返回
- 实现其实和 LastNode 是一样的。只不过换个条件跳出去就行了
- 这个就没什么详细注释的，前面的懂了这个也就理解了

```go
func (s *SingleList) NodeWithValue(property int) *Node {
	node := new(Node)
	for node = s.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			break
		}
	}
	return node
}
```

### AddAfter 指定某个节点后插入

- 理一下思路，方法会接受 2 个参数，第一个参数用来找到我们要在哪个节点后插入，
- 第二个参数是我们生成的新节点要添加的节点，首先根据第一个参数找到指定的节点，
- 然后这个节点指向下一个节点的指针赋值给新节点的 nextNode 这样这个新节点就会指向
- 指定节点的下一个指向，然后将这个新节点在赋值给指定节点的下一个节点就好了。

```GO
func (s *SingleList) AddAfter(nodeProperty, property int) {
  // 生成一个新节点
	node := &Node{
		property: property,
	}
  // 根据上面定义的这个NodeWithValue方法返回
  // 指定的节点
	specialNode := s.NodeWithValue(nodeProperty)
	if specialNode != nil {
    // 把这个指定节点的下一个指向赋值给新节点
    // 让这个新节点去指向
		node.nextNode = specialNode.nextNode
    // 然后将这个指定的节点下一个指向
    // 指向新的节点
		specialNode.nextNode = node
	}
}
```

### RemoveFromEnd 从最后删除节点

- 思路和找到最后一个节点是一样的，无非是判断的条件变一下
- 找到最后一个节点的前一个节点，让后将这个节点的 nextNode 设置成 nil 就可以了

```GO
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
```

### 翻转单链表

- 翻转整个单链表有两种实现一种是遍历的方式时间复杂度是 O(n) 空间复杂度是 O(1)
- 第二种是递归的方式比较耗费空间 时间复杂度是 O(n) 空间复杂度是 O(n)

#### 遍历实现

- 遍历实现翻转的思路。定义个新链表的头结点，先把当前结点的 next 指针保留，不然更改了指向就找不到了
- 将当前头结点的 next 指向变成指向我们的新的头结点。将更改后的当前结点赋值给新的头结点。
- 将当前结点的值替换成保留的 next。进入下一次的循环。这样一直遍历到最后

```GO
func (s *SingleList) Reverse() {
	var pre *Node
	cur := s.headNode
  // 判断当前链表是空的还是只有一个
  // 这种情况反不反转都是自己直接return
	if s.headNode == nil || s.headNode.nextNode == nil {
		return
	}
	for cur != nil {
    // 保留当前结点的nextNode指针
		next := cur.nextNode
    // 更改当前结点的nextNode指向新的pre结点
		cur.nextNode = pre
    // 将这个更改nextNode的cur当前结点赋值给pre
		pre = cur
    // 将next保留的下一个节点赋值给cur 进入下一次循环
		cur = next
	}
  // 将翻转后的新pre头结点赋值给headNode
	s.headNode = pre
}
```

#### 递归实现

