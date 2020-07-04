package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func main() {
	for i := 0; i < 10; i++ {
		go programmer()
	}
	for i := 0; i < 5; i++ {
		go tester()
	}
	select {}
}
