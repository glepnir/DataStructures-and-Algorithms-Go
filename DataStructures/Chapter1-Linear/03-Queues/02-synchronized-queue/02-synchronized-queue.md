## Synchronized Queue 同步队列

假设我们在一家 IT 公司上班,为了让员工有机会互相了解和放松，公司购买了一个兵乓球桌但是有一些规则：

- 只能单打不能双打
- 两个人打完一局就切换下一对
- 测试人员只能与程序员一起工作，反之亦然(不能有两个测试人员或两个程序员在一起)。 如果程序员或测试员
  想要玩，则需要分别等待测试员或程序员建立有效的配对。

```GO
func main() {
    for i := 0; i < 10; i++ {
        go programmer()
    }
    for i := 0; i < 5; i++ {
        go tester()
    }
    select {} // long day at work...
}
func programmer() {
    for {
        code()
        pingPong()
    }
}
func tester() {
    for {
        test()
        pingPong()
    }
}
```

简单的进行测试，使用 time.sleep 模拟每局游戏的时间。

```GO
func test() {
    work()
}
func code() {
    work()
}
func work() {
    // Sleep up to 10 seconds.
    time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
}
func pingPong() {
    // Sleep up to 2 seconds.
    time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
}
func programmer() {
    for {
        code()
        fmt.Println("Programmer starts")
        pingPong()
        fmt.Println("Programmer ends")
    }
}
func tester() {
    for {
        test()
        fmt.Println("Tester starts")
        pingPong()
        fmt.Println("Tester ends")
    }
}
```

运行程序员会得到差不多这样的输出

```log
>>>> go run main.go
Programmer starts
Tester starts
Programmer starts
Programmer ends
Programmer starts
Tester ends
Tester starts
Tester starts
Tester starts
Programmer ends
Tester ends
Programmer ends
Tester starts
Tester ends
Tester ends
Programmer starts
Tester ends
Programmer starts
Tester starts
Programmer starts
```

但是按照我们的游戏规则，我们应该得到的输出是程序员与测试员是结对开始和结束的。

```
Tester starts
Programmer starts
Tester ends
Programmer ends

Tester starts
Programmer starts
Programmer ends
Tester ends

Programmer starts
Tester starts
Tester ends
Programmer ends

Programmer starts
Tester starts
Programmer ends
Tester ends
```

所以，首先，要么是测试员，要么是程序员走近桌子。 之后，合作伙伴加入(相应的程序员或测试人员)。 当他们
离开游戏时，他们可以按任何顺序进行。 这就是为什么我们有 4 个有效序列。

下面是两个解决方案。 第一种是基于互斥锁的，第二种是使用单独的 Worker Goroutine，它协调整个过程，确保
所有事情都按照游戏规则进行。

## 解决方案 1

这两种解决方案都使用队列数据结构，在接近工作台之前对程序员和测试人员进行排队。 当至少有一个有效的配对
可用时(Dev+QA)，则允许一个配对打乒乓球。

```go
func tester(q *queue.Queue) {
    for {
        test()
        q.StartT()
        fmt.Println("Tester starts")
        pingPong()
        fmt.Println("Tester ends")
        q.EndT()
    }
}
func programmer(q *queue.Queue) {
    for {
        code()
        q.StartP()
        fmt.Println("Programmer starts")
        pingPong()
        fmt.Println("Programmer ends")
        q.EndP()
    }
}
func main() {
    q := queue.New()
    for i := 0; i < 10; i++ {
        go programmer(q)
    }
    for i := 0; i < 5; i++ {
        go tester(q)
    }
    select {}
}
```

队列包(package queue):

```go
package queue
import "sync"
type Queue struct {
    mut                   sync.Mutex
    numP, numT            int
    queueP, queueT, doneP chan int
}
func New() *Queue {
    q := Queue{
        queueP: make(chan int),
        queueT: make(chan int),
        doneP:  make(chan int),
    }
    return &q
}
func (q *Queue) StartT() {
    q.mut.Lock()
    if q.numP > 0 {
        q.numP -= 1
        q.queueP <- 1
    } else {
        q.numT += 1
        q.mut.Unlock()
        <-q.queueT
    }
}
func (q *Queue) EndT() {
    <-q.doneP
    q.mut.Unlock()
}
func (q *Queue) StartP() {
    q.mut.Lock()
    if q.numT > 0 {
        q.numT -= 1
        q.queueT <- 1
    } else {
        q.numP += 1
        q.mut.Unlock()
        <-q.queueP
    }
}
func (q *Queue) EndP() {
    q.doneP <- 1
}
```

队列具有用于两个目的的互斥体：

- 同步对共享计数器(numT 和 numP)的访问
- 充当扮演员工持有的令牌，阻止其他人加入桌子

程序员和测试人员正在等待使用无缓冲信道的乒乓球搭档

```
<-q.queueP 或者 q<-queueT
```
