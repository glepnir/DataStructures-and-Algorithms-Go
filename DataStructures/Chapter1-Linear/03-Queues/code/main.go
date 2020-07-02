package main

import (
	"container/list"
	"fmt"
)

// 入队
func enqueue(queue []int, element int) []int {
	// append 添加至切片的尾部
	queue = append(queue, element)
	fmt.Println("Enqueued:", element)
	return queue
}

// 出队
func dequeue(queue []int) []int {
	// 第一个元素出队
	element := queue[0]
	fmt.Println("Dequeued:", element)
	// 在原slice上reslice,返回除第一个元素后的切片
	return queue[1:]
}

func ListQueue() {
	queue := list.New()

	// pushBack方法会添加到链表的尾部
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)

	// Front会得到当前链表的头部
	front := queue.Front()
	fmt.Println(front.Value)
	// remove将删除分配的内存并避免内存泄漏
	queue.Remove(front)
}

func main() {
	var queue []int // Make a queue of ints.

	queue = enqueue(queue, 10)
	queue = enqueue(queue, 20)
	queue = enqueue(queue, 30)

	fmt.Println("Queue:", queue)

	queue = dequeue(queue)
	fmt.Println("Queue:", queue)

	queue = enqueue(queue, 40)
	fmt.Println("Queue:", queue)
}
