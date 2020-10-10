// Package main provides ...
package main

import (
	"fmt"
	"strconv"
)

type Element struct {
	elementvalue int
}

func (element *Element) String() string {
	return strconv.Itoa(element.elementvalue)
}

type Stack struct {
	elements     []*Element
	elementCount int
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	stack.elementCount++
}

func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}
	stack.elementCount--
	return stack.elements[stack.elementCount]
}

// main method
func main() {
	stack := NewStack()
	stack.Push(&Element{3})
	stack.Push(&Element{5})
	stack.Push(&Element{7})
	stack.Push(&Element{9})
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}
