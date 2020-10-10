## 冒泡排序算法

冒泡排序算法是一种排序算法，它比较一对相邻元素，如果它们的顺序错误，则交换它们。 该算法的复杂度为O(N2)
其中n是要排序的元素的数量。 最小或最大的值向上冒泡到集合的顶部，或者最小或最大的下沉到底部(取决于您
是按升序还是降序排序)。下面的代码片段显示了冒泡排序算法的实现。 BubbleSorter函数接受一个整数数组，并
按升序对该数组的元素进行排序。

## 示例

Main方法初始化一个整形的数组并调用bubbleSorter函数，如下所示:

```go
//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and bytes package
import (
	"fmt"
)

//bubble Sorter method
func bubbleSorter(integers [11]int) {
	var num int
	num = 11
	var isSwapped bool
	isSwapped = true
	for isSwapped {
		isSwapped = false
		var i int
		for i = 1; i < num; i++ {
			if integers[i-1] > integers[i] {
				var temp = integers[i]
				integers[i] = integers[i-1]
				integers[i-1] = temp
				isSwapped = true
			}
		}
	}
	fmt.Println(integers)
}

// main method
func main() {
	var integers [11]int = [11]int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("Bubble Sorter")
	bubbleSorter(integers)
}

```
