## 选择排序

选择排序是一种将输入集合分成两个片段的算法。 通过将列表中最小或最大的元素从列表的左侧交换到右侧来对

元素的子列表进行排序。 该算法的阶数为 O(n2)。 该算法在处理大集合时效率较低，且性能比插入排序算法差。

下面的代码显示了SelectionSorter函数的实现，该函数接受要排序的集合：

```go
// 选择排序方法
func SelectionSorter(elements []int) {
	for i := 0; i < len(elements)-1; i++ {
		min := i
		for j := i + 1; j <= len(elements)-1; j++ {
			if elements[j] < elements[min] {
				min = j
			}
		}
		swap(elements, i, min)
	}
}
```

main函数:

```go
// swap method
func swap(elements []int, i int, j int) {
	var temp int
	temp = elements[j]
	elements[j] = elements[i]
	elements[i] = temp
}

//main method
func main() {
	var elements []int
	elements = []int{11, 4, 18, 6, 19, 21, 71, 13, 15, 2}
	fmt.Println("Before Sorting ", elements)
	SelectionSorter(elements)
	fmt.Println("After Sorting", elements)
}
```

`go run sorting_selection.go` 输出:

```
Before Sorting  [11 4 18 6 19 21 71 13 15 2]
After Sorting [2 4 6 11 13 15 18 19 21 71]
```
