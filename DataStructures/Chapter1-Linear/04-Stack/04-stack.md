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

```LOG
$ go test -test.bench=".*" -benchmem -v
goos: darwin
goarch: amd64
pkg: test/test8
Benchmark_Push-4   	10000000	         222 ns/op	      94 B/op	       0 allocs/op
Benchmark_Pop-4    	20000000	        65.0 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	test/test8	7.644s
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
pkg: test/test10
Benchmark_Push-4   	 5000000	       222 ns/op	      48 B/op	       1 allocs/op
Benchmark_Pop-4    	20000000	        73.5 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	test/test10	10.837s
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
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	n := &node{value, this.top}
	this.top = n
	this.length++
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
pkg: test/test9
Benchmark_Push-4   	10000000	         178 ns/op	      32 B/op	       1 allocs/op
Benchmark_Pop-4    	20000000	        75.5 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	test/test9	9.776s
```

### 对比

三种方式，总的来看，第三种基于自定义数据结构的实现方式，在 push 上效率最高，而且实现也比较精巧。个人其
实是推荐使用这种方式的。其次，是基于 container/list 实现的方式。

```
特性对比	push速度	pop速度	push内存分配	pop内存分配
基于slice	222ns/op	65ns/op	94B/op	0B/op
container/list链表	222ns/op	73.5ns/op	48B/op	0B/op
自定义数据结构	178ns/op	75ns/op	32B/op	0B/op
```

