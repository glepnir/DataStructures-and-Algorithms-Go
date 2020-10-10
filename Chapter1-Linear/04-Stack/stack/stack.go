package main

import (
	"fmt"
	"sync"
)

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
