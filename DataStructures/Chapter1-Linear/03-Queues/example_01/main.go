//
// Copyright (C) 2020 Raphael <sawakitaeiji233@gmail.com>
//
// DataStructuresAndAlgorithms - Go : Queue
//
//

package main

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
