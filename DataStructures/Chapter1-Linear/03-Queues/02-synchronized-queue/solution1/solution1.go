// Package main provides ...
package solution1

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

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
	q := New()
	for i := 0; i < 10; i++ {
		go programmer(q)
	}
	for i := 0; i < 5; i++ {
		go tester(q)
	}
	select {}
}
