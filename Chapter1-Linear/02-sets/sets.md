# Set 集合

Set 是线性数据结构，特点在于 Set 的 item 是不能重复的。那么如何在 Go 中实现？联想到 Go 的 map，当我们在 map 中
定义重复的 key 时，会得到错误提示重复的 key。这样就比较符合 Set 的特性了。

## 实现

我们可以通过定义一个 map key 来存储我们的值，value 是 bool 类型来判断是否存在。一个空的 Set 结构会使这样的

``` GO
// 空的Set结构
set := make(map[string]bool)
// 进行赋值
set["Foo"] = true
// 遍历Set结构
for k := range set {
    fmt.Println(k)
}
// 删除
delete(set, "Foo")
// 获得Set结构的大小也就是长度
size := len(set)
// 通过获取value 的bool值判断是否存在
exists := set["Foo"]
```

通过上面的代码使用 map 很简单的实现了 Set 结构。但是 value 的 bool 类型会使用内存。这里要引出 go 中
一个比较特殊的知识点空结构体，空间结构体并不占用内存，经常能看到空间结构体配合 channel 使用。

## 验证空结构体

* 可以使用 unsafe 包。 `unsafe.Sizeof` 会返回参数所占据的字节数。

``` go
// 定义非空结构体
type S struct {
        a uint16
        b uint32
}

func main() {
	var s S
	fmt.Println(unsafe.Sizeof(s)) // 输出8
	var s2 struct{}
	fmt.Println(unsafe.Sizeof(s2)) // 输出0
}
```

由此可见空结构体的内存占用是 0 字节。

* 那么 2 个空结构体的内存地址是多少呢？

``` go
a := struct{}{}
b := struct{}{}
fmt.Println(a == b) // 输出true
fmt.Printf("%p, %p\n", &a, &b) // 相同的内存地址
```

通过输出也看到它们是指向相同的内存地址。

## 优化

``` go
type void struct{}
var member void

set := make(map[string]void)
set["Foo"] = member
for k := range set {
    fmt.Println(k)
}
delete(set, "Foo")
size := len(set)
_, exists := set["Foo"]
```

## 实现常用的方法

接下来实现一些常用的方法。

### Add 添加方法

* 通过可变长的参数我们可以添加多个值到 Set 中，实现也很简单，通过 for range 可变长的参数切片，将参数
* 赋值到 map 的 key 中。value 的值赋予 Exists 空结构体。

``` go
var Exists = struct{}{}

type Set struct {
	m map[interface{}]struct{}
}

func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.m[item] = Exists
	}
	return nil
}
```

### New 简单工厂初始化

* 结合 Add 方法。构造一个简单工厂函数 New 返回一个初始化的 Set

``` go
func New(items ...interface{}) *Set {
    // 获取Set的地址
	s := &Set{}
	// 声明map类型的数据结构
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}
```

### Contains 包含

* Contains 操作其实就是查询操作，看看有没有对应的 Item 存在，可以利用 Map 的特性来实现，但是由于不需
* 要 Value 的数值，所以可以用 `_,ok` 来达到目的。

``` go
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}
```

### Size

* 返回 Set 结构体的长度

``` go
func (s *Set) Size() int {
	return len(s.m)
}
```

### Clear 清除

* 清除操作可以通过重新初始化一个 Set 来实现。

``` go
func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}
```

### Equal 相等

* 判断两个 Set 是否相等，可以通过循环遍历来实现，即将 A 中的每一个元素，查询在 B 中是否存在，只要有一
* 个不存在，A 和 B 就不相等，实现方式如下所示：

``` GO
func (s *Set) Equal(other *Set) bool {
  // 如果两者Size不相等，就不用比较了
	if s.Size() != other.Size() {
		return false
	}

   // 迭代查询遍历
	for key := range s.m {
        // 只要有一个不存在就返回false
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
```

### IsSubSet 子集

* 判断 A 是不是 B 的子集，也是循环遍历的过程，具体分析在上面已经讲述过，实现方式如下所示：

``` GO
func (s *Set) IsSubset(other *Set) bool {
	// s的size长于other，不用说了
	if s.Size() > other.Size() {
		return false
	}
    // 迭代遍历
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
```
