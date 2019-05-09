// Package main provides ...
package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int
	front   int //指向队首前面一个不包含队首
	rear    int
}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add添加数据到队列")
		fmt.Println("2.输入get从队列中获取数据")
		fmt.Println("3.输入show显示队列")
		fmt.Println("4.输入exit退出队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要添加的队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("添加队列成功")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("取出了一个数", val)
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}

//ADD data to Queue
func (q *Queue) AddQueue(val int) error {
	if q.rear == q.maxSize-1 {
		return errors.New("queue is full")
	}
	q.rear++
	q.array[q.rear] = val
	return nil
}
func (q *Queue) ShowQueue() {
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, q.array[i])
	}
	fmt.Println()
}

func (q *Queue) GetQueue() (value int, err error) {
	if q.front == q.rear {
		return -1, errors.New("queue is empty")
	}
	q.front++
	value = q.array[q.front]
	return value, nil
}
