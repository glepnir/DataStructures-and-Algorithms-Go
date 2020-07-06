## Stack 栈

堆栈是后进先出(LIFO)的结构，其中项目是从顶部添加的。 在解析器中使用堆栈来求解迷宫算法。 PUSH、POP、TOP 和
GET SIZE 是堆栈数据结构上允许的典型操作。 语法分析、回溯和编译时内存管理是可以使用堆栈的一些真实场景
堆栈实现示例如下(stack.go)：

### 实现

```go
// 栈内元素结构体
type Element struct {
	elementvalue int
}

func (element *Element) String() string {
	return strconv.Itoa(element.elementvalue)
}

type Stack struct {
  // 使用切片存储元素
	elements     []*Element
  // 栈元素计数
	elementCount int
}
```

使用一个简单工厂函数返回一个初始化的栈

```GO
func NewStack() *Stack {
	return &Stack{}
}
```

### Push 入栈方法

Push 方法将节点添加到 Stack 的顶部。 在下面的代码示例中，Stack 的 Push 方法将元素添加到 Elements 切片
并增加 Count 元素，而 Append 方法将元素添加到 Stack：

```GO
func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	stack.elementCount++
}
```

### Pop 出栈方法

Stack 的 Pop 方法从 Element 切片中移除最后一个元素并返回该元素，如下面的代码所示。 len 方法返回元素
切片的长度：

```GO
func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}
	stack.elementCount--
	return stack.elements[stack.elementCount]
}
```

- 输出

```go
// main method
func main() {
	stack := NewStack()
	stack.Push(&Element{3})
	stack.Push(&Element{5})
	stack.Push(&Element{7})
	stack.Push(&Element{9})
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}

--- output
9 7 5 3
```
