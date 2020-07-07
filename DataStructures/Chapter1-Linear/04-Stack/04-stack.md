## Stack 栈

堆栈是后进先出(LIFO)的结构，其中项目是从顶部添加的。 在解析器中使用堆栈来求解迷宫算法。 PUSH、POP、TOP 和
GET SIZE 是堆栈数据结构上允许的典型操作。 语法分析、回溯和编译时内存管理是可以使用堆栈的一些真实场景
堆栈实现示例如下(stack.go)：

栈实现常用的三种经典方式使用切片 链表 自定义结构体。

### 切片实现

```GO
// Item the type of the stack
type Item interface{}

// ItemStack the stack of Items
type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

// New creates a new ItemStack
func NewStack() *ItemStack {
	s := &ItemStack{}
	s.items = []Item{}
	return s
}

// Pirnt prints all the elements
func (s *ItemStack) Print() {
	fmt.Println(s.items)
}

// Push adds an Item to the top of the stack
func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	s.lock.Unlock()
	s.items = append(s.items, t)
}

// Pop removes an Item from the top of the stack
func (s *ItemStack) Pop() Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}
```

使用切片的特点:

- 利用 Golang 的 Slice，足够简单
- 允许添加任意类型的元素
- Push 和 Pop 有加锁处理，线程安全
- 使用 slice 作为 Stack 的底层支持，速度更快。但是，slice 最大的问题其实是存在一个共用底层数组的的问题，哪
  怕你不断的 Pop，但可能对于 Golang 来说，其占用的内存，并不一定减少。

性能测试如下：

```GO
package main

import (
	"testing"
)

var stack *ItemStack

func init() {
	stack = NewStack()
}

func Benchmark_Push(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		stack.Push("test")
	}
}

func Benchmark_Pop(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		stack.Push("test")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		stack.Pop()
	}
}
```

> 以下所有测试基于 go 1.14.4 darwin/amd64

```LOG
$ go test -test.bench=".*" -benchmem -v
goos: darwin
goarch: amd64
pkg: github.com/glepnir/DataStructuresAndAlgorithms-Go/DataStructures/Chapter1-Linear/04-Stack/stack
Benchmark_Push
Benchmark_Push-8   	12399790	        89.7 ns/op	      97 B/op	       0 allocs/op
Benchmark_Pop
Benchmark_Pop-8    	39666664	        29.4 ns/op	       0 B/op	       0 allocs/op
PASS
```

### 链表实现

```GO
package main

import (
	"container/list"
	"sync"
)

type Stack struct {
	list *list.List
	lock *sync.RWMutex
}

func NewStack() *Stack {
	list := list.New()
	l := &sync.RWMutex{}
	return &Stack{list, l}
}

func (stack *Stack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}
```

简单来说，container/list 是一个链表。用链表模拟栈，要么都向链表的最后做 push 和 pop，要么都向链表的起点
做 push 和 pop，这种方法也非常简单。

性能测试如下：

```LOG
$ go test -test.bench=".*" -benchmem -v  -count=1
goos: darwin
goarch: amd64
pkg: github.com/glepnir/DataStructuresAndAlgorithms-Go/DataStructures/Chapter1-Linear/04-Stack/stack2
Benchmark_Push
Benchmark_Push-8   	 8214225	       158 ns/op	      48 B/op	       1 allocs/op
Benchmark_Pop
Benchmark_Pop-8    	32743381	        43.3 ns/op	       0 B/op	       0 allocs/op
PASS
```

### godoc 的实现参考（自定义数据结构实现）

```go
package main

import (
	"sync"
)

type (
	Stack struct {
		top    *node
		length int
		lock   *sync.RWMutex
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func NewStack() *Stack {
	return &Stack{nil, 0, &sync.RWMutex{}}
}

// Return the number of items in the stack
func (stack *Stack) Len() int {
	return stack.length
}

// View the top item on the stack
func (stack *Stack) Peek() interface{} {
	if stack.length == 0 {
		return nil
	}
	return stack.top.value
}

// Pop the top item of the stack and return it
func (stack *Stack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.length == 0 {
		return nil
	}
	n := stack.top
	stack.top = n.prev
	stack.length--
	return n.value
}

// Push a value onto the top of the stack
func (stack *Stack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	n := &node{value, stack.top}
	stack.top = n
	stack.length++
}
```

此方式特点:

- 实现比较精巧，唯一需要理解的地方就是 node 结构体中，prev 的部分，它指向的是前一个 node 成员。
- 允许添加任意类型的元素。
- Push 和 Pop 是线程安全的。

性能测试如下：

```GO
$ go test -test.bench=".*" -benchmem -v
goos: darwin
goarch: amd64
pkg: github.com/glepnir/DataStructuresAndAlgorithms-Go/DataStructures/Chapter1-Linear/04-Stack/stack3
Benchmark_Push
Benchmark_Push-8   	10694956	        99.9 ns/op	      32 B/op	       1 allocs/op
Benchmark_Pop
Benchmark_Pop-8    	38457697	        30.2 ns/op	       0 B/op	       0 allocs/op
PASS
```

### 对比

| 特性对比            | push 速度 | pop 速度  | push 内存分配 | pop 内存分配 | allocs(不同的内存分配) |
| ------------------- | --------- | --------- | ------------- | ------------ | ---------------------- |
| 基于 slice          | 89.7ns/op | 29.4ns/op | 97B/op        | 0B/op        | 0 allocs/op            |
| container/list 链表 | 158ns/op  | 43.3ns/op | 48B/op        | 0B/op        | 1 allocs/op            |
| 自定义数据结构      | 99.9ns/op | 30.2ns/op | 32B/op        | 0B/op        | 1 allocs/op            |
