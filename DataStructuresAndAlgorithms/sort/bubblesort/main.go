// Package main provides ...
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{24, 69, 80, 57, 13}
	bubblesort(s)
	fmt.Println(s)
	data := sort.IntSlice{22, 34, 3, 40, 18, 4}
	BubbleSort(data)
	fmt.Println(data)
}

func bubblesort(arr []int) []int {
	var tmp int
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			fmt.Println(j)
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

// BubbleSort 冒泡排序. data必须实现sort包中的Interface接口
func BubbleSort(data sort.Interface) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		isChanged := false
		for j := 0; j < n-1-i; j++ {
			if data.Less(j, j+1) {
				data.Swap(j, j+1)
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}
