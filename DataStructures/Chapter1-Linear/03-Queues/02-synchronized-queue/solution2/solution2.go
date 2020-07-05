// Package solution2 provides ...
package solution2

import (
	"fmt"
	"math/rand"
	"time"
)

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

func tester(q *Queue) {
	for {
		work()
		q.StartT()
		fmt.Println("Tester starts")
		pingPong()
		fmt.Println("Tester ends")
		q.EndT()
	}
}
func programmer(q *Queue) {
	for {
		work()
		q.StartP()
		fmt.Println("Programmer starts")
		pingPong()
		fmt.Println("Programmer ends")
		q.EndP()
	}
}

func work() {
	// Sleep up to 10 seconds.
	time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
}
func pingPong() {
	// Sleep up to 2 seconds.
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
}

func main() {
	queue := New()
	for i := 0; i < 10; i++ {
		go programmer(queue)
	}
	for i := 0; i < 5; i++ {
		go tester(queue)
	}
	select {}
}
