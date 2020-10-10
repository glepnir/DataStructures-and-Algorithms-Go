// Package main provides ...
package main

import "testing"

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
