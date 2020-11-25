# 合并排序算法

合并排序算法是约翰·冯·诺伊曼(John Von Neumann)发明的一种基于比较的方法。 比较相

邻列表中的每个元素以进行排序。 该算法的性能为O(nlogn)量级。 此算法是对链表进行

排序的最佳算法。

下面的代码片段演示了合并排序算法。 CreateArray函数将num int作为参数，并返回一个

由随机化元素组成的整数切片：

## 实现

下面的代码片段演示了合并排序算法。 CreateArray函数将num int作为参数，并返回一个由随机化元素组成的整

数ARRAY：

``` go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// create array
func createArray(num int) []int {
	var array []int
	array = make([]int, num, num)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		rand.Intn(99999) - rand.Intn(99999)
	}
	return array
}
```

`MergeSorter` 方法接受整数元素切片作为参数，两个元素子切片递归传递给MergeSorter方法。 结果切片被联接并

作为集合返回，如下所示：

``` go
// MergeSorter algorithm
func MergeSorter(array []int) []int {
	if len(array) < 2 {
		return array
	}
	var middle int
	middle = (len(array)) / 2
	return JoinArrays(MergeSorter(array[:middle]), MergeSorter(array[middle:]))
}
```

JoinArray函数将leftArr和rightArr整数切片作为参数。 组合后的切片在以下代码中返回：

``` go
func JoinArrays(leftArr []int, rightArr []int) []int {
	num, i, j := len(leftArr)+len(rightArr), 0, 0
	array := make([]int, num, num)
	for k := 0; k < num; k++ {
		if i > len(leftArr)-1 && j <= len(rightArr)-1 {
			array[k] = rightArr[j]
			j++
		} else if j > len(rightArr)-1 && i <= len(leftArr)-1 {
			array[k] = leftArr[i]
			i++
		} else if leftArr[i] < rightArr[j] {
			array[k] = leftArr[i]
			i++
		} else {
			array[k] = rightArr[j]
			j++
		}
	}
	return array
}
```

Main方法对包含40个元素的整数切片进行初始化，并在排序前后打印元素，如下所示：

``` go
func main() {
	var elements []int
	elements = createArray(40)
	fmt.Println("\n Before Sorting \n\n", elements)
	fmt.Println("\n-After Sorting\n\n", MergeSorter(elements), "\n")
}
```
