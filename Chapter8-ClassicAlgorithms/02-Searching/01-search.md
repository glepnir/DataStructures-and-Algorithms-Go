# 线性搜索

线性搜索方法通过顺序检查集合中的每个元素来找到集合中的给定值。 线性搜索算法的时间复杂度为O（n）。二

进制搜索算法和哈希表的性能优于此搜索算法。以下代码段显示了线性搜索方法的实现。  LinearSearch函数采用

整数元素数组和findElement int作为参数。 如果找到findElement，则该函数返回一个布尔值true；否则，返回true。 否则，返回false：

``` go
package main

import {
  fmt
}

func LinearSearch(elements []int,findElement int) bool {
  for _,element = range elements{
    if element == findElement {
      return true
    }
  }
  return false
}
```

## 二进制搜索(二分查找)

二进制搜索算法将输入值与已排序集合的中间元素进行比较。 如果值不相等，则消除其中找不到元素的一半。 搜

索继续在集合的其余一半上进行。 该算法的时间复杂度约为O（log n）。下面的代码片段显示了使用sort包中的

`sort.Search` 函数执行二进制搜索算法的过程。main方法初始化elements数组并调用sort. Search函数以找到一个整数元素：

``` go
package main

import {
  "fmt"
  "sort"
}

fun main() {
  elements = []int{1,3,16,10,45,31,28,36,45,75}
  element = 36
  i := sort.Search(len(elements),func(i int) bool {
    return elements[i] >= element
  })
  if i < len(elements) && elements[i] == element{
    fmt.Printf("found element %d at index %d %v\n",element,i,elements)
  }else{
    fmt.Printf("element %d not found in %v\n",element,elements)
  }
}

```

另一种方法使用递归的方式并且不使用sort包:

``` go
func binarySearch(array []int, target int, lowIndex int, highIndex int) int {
    //specify condition to end the recursion
    if highIndex < lowIndex {
        return -1
    }
    // Define our middle index
    mid := int((lowIndex + highIndex) / 2)
    if array[mid] > target {
        return binarySearch(array, target, lowIndex,mid)
    }else if array[mid] < target {
        return binarySearch(array, target,mid+1,highIndex)
    }else {
        return mid
    }
}
```

## 插值搜索

插值搜索算法在排序的集合中搜索元素。 该算法通过减少估计位置之前或之后的搜索空间来在估计位置找到输入元

素。 搜索算法的时间复杂度约为O（log log n）。以下代码片段实现了插值搜索算法。  InterpolationSearch函

数将整数元素的数组和要找到的整数元素作为参数。 该函数在集合中查找元素，并返回布尔值和找到的元素的索引：
