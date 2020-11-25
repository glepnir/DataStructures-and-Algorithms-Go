# 插入排序

插入排序是一种一次创建一个元素的最终排序数组的算法。 该算法的性能为O(N2)量级。 与其他算法(如快速、堆

和合并排序)相比，此算法在大型集合上的效率较低。 在现实生活中，插入排序的一个很好的例子是棋牌游戏中玩

家手动排序卡片的方式。插入排序算法的实现如以下代码片段所示。 RandomSequence函数将元素数量作为参数，

并返回随机整数数组：



``` go
 // randomSequence method
func randomSequence(num int) []int {
	var sequence []int
	sequence = make([]int, num, num)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		sequence[i] = rand.Intn(999) - rand.Intn(999)
	}
	return sequence
}
```

### 插入排序实现

InsertionSorter方法的实现如以下代码片断所示。 此方法将整数数组作为参数并对其进行排序：

``` go
//InsertionSorter method
func InsertionSorter(elements []int) {
	var n = len(elements)
	var i int
	for i = 1; i < n; i++ {
		var j int
		j = i
		for j > 0 {
			if elements[j-1] > elements[j] {
				elements[j-1], elements[j] = elements[j], elements[j-1]
			}
			j = j - 1
		}
	}
}
```

Main方法通过调用随机序列函数来初始化序列，如以下代码所示。 InsertionSorter函数获取序列并按升序进行排序：

``` go
//main method
func main() {
	var sequence []int
	sequence = randomSequence(24)
	fmt.Println("\n^^^^^^ Before Sorting ^^^ \n\n", sequence)
	InsertionSorter(sequence)
	fmt.Println("\n--- After Sorting ---\n", sequence)
}
```

`go run sorting_insertion.go` 查看输出

```
^^^^^^ Before Sorting ^^^

 [-86 -816 -205 86 -106 -624 357 189 -263 149 419 125 -277 -72 235 117 -361 -321 528 -170 834 85 519 -89]

--- After Sorting ---
 [-816 -624 -361 -321 -277 -263 -205 -170 -106 -89 -86 -72 85 86 117 125 149 189 235 357 419 519 528 834]
```
