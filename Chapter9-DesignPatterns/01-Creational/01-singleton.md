## 单例模式(Singleton design pattern) -在整个程序中具有类型的唯一实例

### 概念

单例模式(Singleton design pattern)很容易记住。顾名思义，它将为您提供一个对象的单个实例，并保证没有重复。

在首次使用实例的调用中，将创建该实例，然后在应用程序中需要使用该特定行为的所有部分之间重用它。您将在

许多不同的情况下使用单例模式。 例如：

- 当您想使用与数据库的相同连接进行每个查询时。

- 当您打开与服务器的安全Shell（SSH）连接以执行一些任务时，并且不想为每个任务重新打开连接。

- 如果您需要限制对某些变量或空间的访问，可以使用Singleton作为此变量的门（我们将在以下章节中看到，无论如何，使用Go在通道中都可以实现此目的）

- 如果需要限制对某些位置的调用数，可以创建一个单例实例，以便在接受的窗口中进行调用

当然还有很多场景可以使用单例模式，以上只是一些举例。

### 实现

我们通常会编写一个静态方法和实例来检索java或C++语言中的单实例。在Go中，我们没有关键字static，但是我们可以通过使用包的

作用域来获得相同的结果。首先，我们创建一个结构，该结构包含我们希望在程序执行期间保证为单例对象的对象：

```go
type singleton struct{
  count int
}

var instance *singleton

func GetInstance() *singleton{
  if instance == nil {
    instance = new(singleton)
  }
  return instance
}

func (s *singleton) AddOne() int{
  s.count++
  return s.count
}
```

通过判断实例`instance`是否为nil来判断是否已经实例化。这段代码看似没什么问题，但是当我们使用goroutine并发的去

访问这个变量时就会存在问题，保护临界资源的一种手段就是加锁了。例如:

```go

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creting Single Instance Now")
            singleInstance = &single{}
        } else {
            fmt.Println("Single Instance already created-1")
        }
    } else {
        fmt.Println("Single Instance already created-2")
    }
    return singleInstance
}

func main() {
    for i := 0; i < 100; i++ {
        go getInstance()
    }
    // Scanln is similar to Scan, but stops scanning at a newline and
    // after the final item there must be a newline or EOF.
    fmt.Scanln()
}
```

这样通过加锁的方式保护临界资源防止多个goroutine访问产生的数据竞争问题。

另外一种办法Go的同步包提供了一个方法 `sync.Once`,这个方法只会执行一次:

```go
var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        once.Do(
            func() {
                fmt.Println("Creting Single Instance Now")
                singleInstance = &single{}
            })
    } else {
        fmt.Println("Single Instance already created-2")
    }
    return singleInstance
}
```

