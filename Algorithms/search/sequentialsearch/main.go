// Package main provides ...
package main

import (
	"fmt"
)

func main() {
	var name string
	s := []string{"apple", "oriange", "banana"}
	fmt.Scanln(&name)
	if find(s, name) {
		fmt.Println("find success")
	} else {
		fmt.Println("find failed")
	}
}

func find(arr []string, name string) bool {
	for i, _ := range arr {
		if arr[i] == name {
			return true
		}
		return false
	}
	return false
}
