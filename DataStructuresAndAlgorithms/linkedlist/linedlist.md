# 链表
* 链表（Linked list）是一种常见的基础数据结构，是一种线性表，但是并不会按线性的顺序存储数据，而是在每一个节点里存到下一个节点的指针(Pointer)。由于不必须按顺序存储，链表在插入的时候可以达到O(1)的复杂度，比另一种线性表顺序表快得多，但是查找一个节点或者访问特定编号的节点则需要O(n)的时间，而顺序表相应的时间复杂度分别是O(logn)和O(1)。

* 使用链表结构可以克服数组链表需要预先知道数据大小的缺点，链表结构可以充分利用计算机内存空间，实现灵活的内存动态管理。但是链表失去了数组随机读取的优点，同时链表由于增加了结点的指针域，空间开销比较大。

* 链表作为一种基础的数据结构可以用来生成其它类型的数据结构。链表通常由一连串节点组成，每个节点包含任意的实例数据（data fields）和一或两个用来指向上一个/或下一个节点的位置的链接（"links"）。链表最明显的好处就是，常规数组排列关联项目的方式可能不同于这些数据项目在记忆体或磁盘上顺序，数据的访问往往要在不同的排列顺序中转换。而链表是一种自我指示数据类型，因为它包含指向另一个相同类型的数据的指针（链接）。链表允许插入和移除表上任意位置上的节点，但是不允许随机存取。链表有很多种不同的类型：单向链表，双向链表以及循环链表。
![](https://github.com/taigacute/IMG/blob/master/Datastruct/linkedlist.png)

## 单向链表
 * 单链表 ：由各个内存结构通过一个 Next 指针链接在一起组成，每一个内存结构都存在后继内存结构【链尾除外】，内存结构由数据域和 Next 指针域组成。
 ![](https://github.com/taigacute/IMG/blob/master/Datastruct/singlelist.png)
 - Data 数据 + Next 指针，组成一个单链表的内存结构 ；
 - 第一个内存结构称为 链头，最后一个内存结构称为 链尾； 
 - 链尾的 Next 指针设置为 NULL [指向空]；
 - 单链表的遍历方向单一【只能从链头一直遍历到链尾】
 * 单链表操作:
  ![](https://github.com/taigacute/IMG/blob/master/Datastruct/singlelistoperation.png)

## 双向链表
  * 双向链表 : 双向链表 ：由各个内存结构通过指针 Next 和指针 Prev 链接在一起组成，每一个内存结构都存在前驱内存结构和后继内存结构【链头没有前驱，链尾没有后继】，内存结构由数据域、Prev 指针域和 Next 指针域组成。
  ![](https://github.com/taigacute/IMG/blob/master/Datastruct/dblist.png)
  - Data 数据 + Next 指针 + Prev 指针，组成一个双向链表的内存结构；
  - 第一个内存结构称为 链头，最后一个内存结构称为 链尾；
  - 链头的 Prev 指针设置为 NULL， 链尾的 Next 指针设置为 NULL；
  - Prev 指向的内存结构称为 前驱， Next 指向的内存结构称为 后继；
  - 双向链表的遍历是双向的，即如果把从链头的 Next 一直到链尾的[NULL] 遍历方向定义为正向，那么从链尾的 Prev 一直到链头 [NULL ]遍历方向就是反向；
  * 双向链表操作
  ![](https://github.com/taigacute/IMG/blob/master/Datastruct/doublelist.png)
## 循环链表
  * 循环链表 ： 单向循环链表 [Circular Linked List] : 由各个内存结构通过一个指针 Next 链接在一起组成，每一个内存结构都存在后继内存结构，内存结构由数据域和 Next 指针域组成。
双向循环链表 [Double Circular Linked List] : 由各个内存结构通过指针 Next 和指针 Prev 链接在一起组成，每一个内存结构都存在前驱内存结构和后继内存结构，内存结构由数据域、Prev 指针域和 Next 指针域组成。
![](https://github.com/taigacute/IMG/blob/master/Datastruct/%E5%BE%AA%E7%8E%AF%E9%93%BE%E8%A1%A8.png)
  - 循环链表分为单向、双向两种；
  - 单向的实现就是在单链表的基础上，把链尾的 Next 指针直接指向链头，形成一个闭环；
  - 双向的实现就是在双向链表的基础上，把链尾的 Next 指针指向链头，再把链头的 Prev 指针指向链尾，形成一个闭环；
  - 循环链表没有链头和链尾的说法，因为是闭环的，所以每一个内存结构都可以充当链头和链尾；
  * 循环链表操作：
  ![](https://github.com/taigacute/IMG/blob/master/Datastruct/%E5%BE%AA%E7%8E%AF%E9%93%BE%E8%A1%A8%E6%93%8D%E4%BD%9C.png)
