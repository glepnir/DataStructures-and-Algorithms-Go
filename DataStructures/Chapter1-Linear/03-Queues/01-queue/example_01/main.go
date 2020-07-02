//
// Copyright (C) 2020 Raphael <sawakitaeiji233@gmail.com>
//
// DataStructuresAndAlgorithms - Go : Queue
//
//

package main

import "fmt"

// 订单结构体
type Order struct {
	// 优先级
	priority int
	// 数量
	quantity int
	// 产品
	product string
	// 客户名称
	customerName string
}

type Queue []*Order

func (o *Order) NewOrder(priority int, quantity int, product string, customerName string) {
	o.priority = priority
	o.quantity = quantity
	o.product = product
	o.customerName = customerName
}

func (q *Queue) Add(order *Order) {
	if len(*q) == 0 {
		*q = append(*q, order)
	} else {
		appended := false
		for key, addedOrder := range *q {
			if order.priority > addedOrder.priority {
				// 其他代码没什么解释的这一行新手可能看不懂解释一下
				// 比较已添加订单和准备添加订单的优先级
				// 如果准备添加的订单优先级比已经添加的订单的优先级高那么
				// 先生成一个新的订单队列切片元素是准备添加的订单Queue{order}
				// 然后将从key开始的这个订单到最后的订单追加到这个新生成的订单队列
				// 也就是这句代码(*q)[key:]... 三点运算符展开这个切片。
				// 再讲这个新的订单队列切片展开追加到原始切片到key的前面
				// 再通俗一点就是在key的这个位置前面插入了一个元素也就是新订单。
				*q = append((*q)[:key], append(Queue{order}, (*q)[key:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*q = append(*q, order)
		}
	}
}

// main method
func main() {

	var queue Queue
	queue = make(Queue, 0)

	var order1 *Order = &Order{}

	var priority1 int = 2
	var quantity1 int = 20
	var product1 string = "Computer"
	var customerName1 string = "Greg White"

	order1.NewOrder(priority1, quantity1, product1, customerName1)

	var order2 *Order = &Order{}

	var priority2 int = 1
	var quantity2 int = 10
	var product2 string = "Monitor"
	var customerName2 string = "John Smith"

	order2.NewOrder(priority2, quantity2, product2, customerName2)

	queue.Add(order1)

	queue.Add(order2)

	var i int
	for i = 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
