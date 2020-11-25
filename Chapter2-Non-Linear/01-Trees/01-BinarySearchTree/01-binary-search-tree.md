# Binary search tree 二叉搜索树

二叉搜索树，也称为二叉查找树、有序二叉树（ordered binary tree）或排序二叉树（sorted binary tree），
是指一棵空树或者具有下列性质的二叉树：

* 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
* 若任意节点的右子树不空，则右子树上所有节点的值均大于或等于它的根节点的值；
* 任意节点的左、右子树也分别为二叉搜索树；

如下图，左侧的就是一棵二叉搜索树; 而右侧的则不是，因为节点 9 在根节点 10 的右子树上，但
是其值却比 10 小.

![BinarySearchTree](/image/binarysearchtree.png)

* 节点的度： 一个节点含有的子树的个数称为节点的度
* 树的度 ： 一棵树中，最大的节点的度称为树的度
* 叶节点或终端节点：度为零的节点
* 非终端节点或分支节点： 度不为零的节点
* 父亲节点或父节点： 若一个节点含有子节点，则这个节点称为其子节点的父节点
* 孩子节点或子节点：一个节点含有的子树的根节点称为该节点的子节点
* 兄弟节点： 具有相同父节点的节点互称为兄弟节点
* 节点的层次： 从根节点开始，根为第一层，根的子节点为第二层，以此类推
* 深度：对于任意节点 n， n 的深度为从根到 n 的唯一路径长，根的深度为 0
* 高度： 对于任意节点 n，n 的高度为从 n 到一片树叶的最长路径厂，所有树叶的高度为 0
* 堂兄弟节点：父节点在同一层的节点互称堂兄节点
* 子孙： 以某节点为根的子树中任一节点都称为该节点的子孙
* 森林：由 m（m>0）颗互不相交的树的集合称为森林

## 实现

接下来实现一个如下图的二叉搜索树并实现一些关于操作二叉搜索树的方法

![Binarysearchtree-1](/image/binarysearchtree-1.png)

操作二叉树的一些常用方法:

* Insert(t) 将 Item t 插入树中
* Search(t) 如果树中存在 Item t，则返回 true
* InOrderTraverse() 使用顺序遍历访问所有节点
* PreOrderTraverse() 通过前序遍历访问所有节点
* PostOrderTraverse() 使用后序遍历访问所有节点
* Min() 返回具有存储在树中的最小值的 Item
* Max() 返回具有最大值的 Item 存储在树中
* Remove(t) 从树上删除项目 t
* String() 打印树的 CLI 可读渲染

三种遍历方式：

* **中序遍历** 我们通过跟随最小的链接来访问所有节点，直到找到最左边的叶子，然后处理该叶子并通过进入

  链接到当前节点的下一个最小键值移动到其他节点。在上图中:
 `1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 11`

* **前序遍历** 在拜访孩子之前，此方法首先拜访节点。在上图中：

 `8 -> 4 -> 2 -> 1 -> 3 -> 6 -> 5 -> 7 ->10 -> 9 -> 11`

* **后序遍历** 我们找到最小的叶子（最左边的叶子），然后处理同级节点，然后处理父节点，然后转到下一个

  子树，然后向上导航父节点。在上图中: `1 -> 3 -> 2 -> 5 -> 7 -> 6 -> 4 -> 9 -> 11 -> 10 -> 8`

二叉搜索树由具有属性或属性的节点组成：

* key 整形
* value 整形
* 树的左右节点 leftNode 和 rightNode

一个树节点：

``` GO
// TreeNode
type TreeNode struct {
    key int
    value int
    leftNode *TreeNode
    rightNode *TreeNode
}
```

二叉搜索树 BinarySearchTree 由 TreeNode 类型的 rootNode 和 sync. RWMutex 类型的 lock 组成。 通过访问
rootNode 左侧和右侧的节点，可以从 rootNode 遍历二叉搜索树：

``` GO
// BinarySearchTree
type BinarySearchTree struct {
  // 根节点
  rootNode *TreeNode
  lock sync.RWMutex
}
```

## Insert 方法

``` GO
// 插入树节点
func InsertTreeNode(rootNode, newTreeNode *TreeNode) {
	// 如果新节点小于根节点放到左边否则放到右边
	if newTreeNode.key < rootNode.key {
		// 如果根节点的左子节点是空
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			// 否则递归调用参数的根节点则换成根节点的左子节点
			InsertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else {
		// 同上
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		} else {
			// 一样的递归调用将根节点换成根节点的右子节点
			InsertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}
}

// 二叉搜索树的插入元素方法
func (bst *BinarySearchTree) Insert(key, value int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	treeNode := &TreeNode{key, value, nil, nil}
	if bst.rootNode == nil {
		bst.rootNode = treeNode
	} else {
		InsertTreeNode(bst.rootNode, treeNode)
	}
}
```

## InOrderTraverseTree 中序遍历

``` GO
func inOrderTraverseTree(treeNode *TreeNode, f func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, f)
		f(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, f)
	}
}

// 中序遍历 左子树-根节点-右子树
func (bst *BinarySearchTree) InOrderTraverseTree(f func(int)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	inOrderTraverseTree(bst.rootNode, f)
}
```

## PreOrderTraverseTree 先序遍历

``` GO
func preOrderTraverseTree(treeNode *TreeNode, f func(int)) {
	if treeNode != nil {
		f(treeNode.value)
		preOrderTraverseTree(treeNode.leftNode, f)
		preOrderTraverseTree(treeNode.rightNode, f)
	}
}

// 先序遍历：根节点 -> 左子树 -> 右子树
func (bst *BinarySearchTree) PreOrderTraverseTree(f func(int)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrderTraverseTree(bst.rootNode, f)
}
```

## PostOrderTraverseTree 后序遍历

``` GO
func postOrderTraverseTree(treeNode *TreeNode, f func(int)) {
	if treeNode != nil {
		postOrderTraverseTree(treeNode.leftNode, f)
		postOrderTraverseTree(treeNode.rightNode, f)
	}
}

// 后序遍历: 左子树-右子树-根节点
func (bst *BinarySearchTree) PostOrderTraverseTree(treeNode *TreeNode, f func(int)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	postOrderTraverseTree(bst.rootNode, f)
}
```

## MinNode 最小节点

``` GO
// 找到最小的节点
func (bst *BinarySearchTree) MinNode() *int {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	treeNode := new(TreeNode)
	treeNode = bst.rootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.leftNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.leftNode
	}
}
```

## MaxNode 最大节点

``` GO
// 找到最大的节点
func (bst *BinarySearchTree) MaxNode() *int {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	treeNode := new(TreeNode)
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.rightNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.rightNode
	}
}
```

## SearchNode 搜索节点

``` GO
func searchNode(treeNode *TreeNode, key int) bool {
	if treeNode == nil {
		return false
	}
	if key < treeNode.key {
		return searchNode(treeNode.leftNode, key)
	}
	if key > treeNode.key {
		return searchNode(treeNode.rightNode, key)
	}
	return true
}

func (bst *BinarySearchTree) SearchNode(key int) bool {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	return searchNode(bst.rootNode, key)
}
```

## RemoveNode 删除节点

``` GO
func removeNode(treeNode *TreeNode, key int) *TreeNode {
	// 当要删除的节点不存在时
	if treeNode == nil {
		return nil
	}
	// 要删除的节点在左侧时
	if key < treeNode.key {
		treeNode.leftNode = removeNode(treeNode.leftNode, key)
		return treeNode
	}
	// 要删除的节点在右侧的时候
	if key > treeNode.key {
		treeNode.rightNode = removeNode(treeNode.rightNode, key)
		return treeNode
	}
	// 判断节点类型 如果是叶子节点直接删除
	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}
	// 当要删除的节点只有右侧节点
	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}
	// 当要删除的节点只有左侧节点
	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}
	// 要删除的节点有 2 个子节点，找到右子树的最左节点，替换当前节点
	leftmostrightNode := new(TreeNode)
	for {
		if leftmostrightNode != nil && leftmostrightNode.leftNode != nil {
			leftmostrightNode = leftmostrightNode.leftNode
		} else {
			break
		}
	}
	// 使用右子树的最左节点替换当前节点，即删除当前节点
	treeNode.key, treeNode.value = leftmostrightNode.key, leftmostrightNode.value
	treeNode.rightNode = removeNode(treeNode.rightNode, treeNode.key)
	return treeNode
}

func (bst *BinarySearchTree) RemoveNode(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	removeNode(bst.rootNode, key)
}
```

## 辅助方法

* 为了能在终端上更直观的打印整个二叉树定义 2 个方法，这里的 String 方法并不需要要显示的调用。当我们
* 使用 fmt. Println 之类的方法的时候会自动调用这个 String 方法

``` GO
func (bst *BinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.rootNode, 0)
	fmt.Println("------------------------------------------------")
}

func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += " "
		}
		format += "---[ "
		level++
		stringify(treeNode.leftNode, level)
		fmt.Printf(format+"%d\n", treeNode.key)
		stringify(treeNode.rightNode, level)
	}
}
```
