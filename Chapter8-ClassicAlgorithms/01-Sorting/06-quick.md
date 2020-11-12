## 快速排序

快速排序是一种以有组织的方式对集合中的元素进行排序的算法。 并行化快速排序比合并排序和堆排序快两到三

倍。 该算法的性能为O(Nlogn)阶。 该算法是二叉树排序算法的空间优化版本。

在下面的代码片段中，实现了快速排序算法。 QuickSorter函数接受一个整数元素数组，上面是int，下面是int作

为参数。 该函数将数组分成几个部分，并以递归方式进行划分和排序：

```go
package main

import (
	"fmt"
)

//Quick Sorter method
func QuickSorter(elements []int, below int, upper int) {
	if below < upper {
		var part int
		part = divideParts(elements, below, upper)
		QuickSorter(elements, below, part-1)
		QuickSorter(elements, part+1, upper)
	}
}
```

DivideParts方法接受一个整数元素数组，上面是int，下面是int作为参数。 该方法按升序对元素进行排序，如以下代码所示：

```go
// divideParts method
func divideParts(elements []int, below int, upper int) int {
	var center int
	center = elements[upper]
	var i int
	i = below
	var j int
	for j = below; j < upper; j++ {
		if elements[j] <= center {
			swap(&elements[i], &elements[j])
			i += 1
		}
	}
	swap(&elements[i], &elements[upper])
	return i
}
```

swap和main函数:

```go
//swap method
func swap(element1 *int, element2 *int) {
	var val int
	val = *element1
	*element1 = *element2
	*element2 = val
}

// main method
func main() {
	var num int
	fmt.Print("Enter Number of Elements: ")
	fmt.Scan(&num)
	var array = make([]int, num)
	var i int
	for i = 0; i < num; i++ {
		fmt.Print("array[", i, "]: ")
		fmt.Scan(&array[i])
	}
	fmt.Print("Elements: ", array, "\n")
	QuickSorter(array, 0, num-1)
	fmt.Print("Sorted Elements: ", array, "\n")
}
```
