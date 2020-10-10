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

如果没有可用的对手，从 channel 中读取的操作会阻止 Goroutine。因为 channel 是阻塞的。

```GO
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
```

如果 numP 大于 0(至少有一个程序员在等待游戏)，则等待的程序员的数量减 1，并且将允许一个等待的程序员加入该
桌`(q.queue <- 1)`。 这里有趣的是，在此路径期间，互斥体不会被释放，因此它将充当独占访问乒乓球桌的令牌。
如果没有正在等待的程序员，则 NumT(正在等待的测试人员的数量)会增加，并在`<-q.queueT` 上阻塞 goroutine。

StartP 基本相同，但由程序员执行。在游戏中，互斥体将被锁定，因此需要由程序员或测试员释放。 若要仅在双方
都完成游戏时释放互斥体，则使用障碍`doneP`：

```go
func (q *Queue) EndT() {
    <-q.doneP
    q.mut.Unlock()
}
func (q *Queue) EndP() {
    q.doneP <- 1
}
```

如果程序员仍在玩，且测试器已完成，则测试器将在`<-q.dononP`上阻塞。 一旦程序员达到`<-q.doneP`，屏障
就会打开，互斥体将被释放，允许这些员工返回工作岗位。如果 TESTER 仍在玩，则程序员将阻塞`q.doneP <-1`,当
TESTER 完成时，它将从屏障`<-q.doneP`读取，这将解锁程序员，并释放互斥锁以释放桌子。这里有趣的是，互斥
总是由测试人员释放，即使测试人员或程序员可能将其锁定。 这也是为什么这个解决方案乍一看可能不那么明显的
部分原因。

## 解决方案 2

```go
package queue
const (
    msgPStart = iota
    msgTStart
    msgPEnd
    msgTEnd
)
type Queue struct {
    waitP, waitT   int
    playP, playT   bool
    queueP, queueT chan int
    msg            chan int
}
func New() *Queue {
    q := Queue{
        msg:    make(chan int),
        queueP: make(chan int),
        queueT: make(chan int),
    }
    go func() {
        for {
            select {
            case n := <-q.msg:
                switch n {
                case msgPStart:
                    q.waitP++
                case msgPEnd:
                    q.playP = false
                case msgTStart:
                    q.waitT++
                case msgTEnd:
                    q.playT = false
                }
                if q.waitP > 0 && q.waitT > 0 && !q.playP && !q.playT {
                    q.playP = true
                    q.playT = true
                    q.waitT--
                    q.waitP--
                    q.queueP <- 1
                    q.queueT <- 1
                }
            }
        }
    }()
    return &q
}
func (q *Queue) StartT() {
    q.msg <- msgTStart
    <-q.queueT
}
func (q *Queue) EndT() {
    q.msg <- msgTEnd
}
func (q *Queue) StartP() {
    q.msg <- msgPStart
    <-q.queueP
}
func (q *Queue) EndP() {
    q.msg <- msgPEnd
}
```

我们有一个中央协调器在单独的 Goroutine 中运行，它协调整个过程。 Scheduler 通过 msg 频道获取想要放松的新员
工的信息，或者某人是否已经不再打乒乓球。 当接收到任何消息时，调度器的状态被更新：

- 等待开发人员或 QA 的数量增加
- 更新正在玩的员工信息

在接收到任何定义的消息后，调度器将检查是否允许另一对开始游戏：

`if q.waitP > 0 && q.waitT > 0 && !q.playP && !q.playT {`

如果是，则相应地更新状态，并且解锁一个测试员和一个程序员。我们不再使用互斥体(如解决方案 1 中那样)来
管理对共享数据的访问，而是使用了一个单独的 Goroutine，它通过通道与外部世界通信。 这给了我们更地道的
Go 程序。

> Don’t communicate by sharing memory, share memory by communicating.
> 不要通过共享内存来通信，相反，应该通过通信来共享内存
