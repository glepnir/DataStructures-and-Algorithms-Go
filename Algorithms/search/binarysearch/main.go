// Package main provides ...
package main

import "fmt"

func main() {
	s := []int{1, 8, 10, 89, 1000, 1234}
	binariysearch(s, 0, len(s)-1, 89)
}

// 使用递归的二分查找 数组或切片必须是有序的才可以使用二分查找
// 自己是否可以实现一个非递归的呢？不要参考资料自己实现以下
func binariysearch(arr []int, leftindex, rightindex, value int) {
	midindex := (leftindex + rightindex) / 2
	if leftindex > rightindex {
		fmt.Println("failed")
		return
	}
	if arr[midindex] > value {
		binariysearch(arr, leftindex, midindex-1, value)
	} else if arr[midindex] < value {
		binariysearch(arr, midindex+1, rightindex, value)
	} else {
		fmt.Println(midindex)
	}
}
