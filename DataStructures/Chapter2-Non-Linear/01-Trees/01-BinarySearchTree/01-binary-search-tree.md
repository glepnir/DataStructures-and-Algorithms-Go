## Binary search tree 二叉搜索树

二叉搜索树，也称为二叉查找树、有序二叉树（ordered binary tree）或排序二叉树（sorted binary tree），
是指一棵空树或者具有下列性质的二叉树：

- 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
- 若任意节点的右子树不空，则右子树上所有节点的值均大于或等于它的根节点的值；
- 任意节点的左、右子树也分别为二叉搜索树；

如下图，左侧的就是一棵二叉搜索树;而右侧的则不是，因为节点 9 在根节点 10 的右子树上，但
是其值却比 10 小.

![BinarySearchTree](/image/binarysearchtree.png)

## 实现

接下来实现一个如下图的二叉搜索树并实现一些关于操作二叉搜索树的方法

![Binarysearchtree-1](/image/binarysearchtree-1.png)

操作二叉树的一些常用方法:

- Insert(t) 将 Item t 插入树中
- Search(t) 如果树中存在 Item t，则返回 true
- InOrderTraverse() 使用顺序遍历访问所有节点
- PreOrderTraverse() 通过前序遍历访问所有节点
- PostOrderTraverse() 使用后序遍历访问所有节点
- Min() 返回具有存储在树中的最小值的 Item
- Max() 返回具有最大值的 Item 存储在树中
- Remove(t) 从树上删除项目 t
- String() 打印树的 CLI 可读渲染

三种遍历方式：

- **中序遍历** 我们通过跟随最小的链接来访问所有节点，直到找到最左边的叶子，然后处理该叶子并通过进入
  链接到当前节点的下一个最小键值移动到其他节点。在上图中:
  `1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 11`

- **前序遍历** 在拜访孩子之前，此方法首先拜访节点。在上图中：
  `8 -> 4 -> 2 -> 1 -> 3 -> 6 -> 5 -> 7 ->10 -> 9 -> 11`

- **后序遍历** 我们找到最小的叶子（最左边的叶子），然后处理同级节点，然后处理父节点，然后转到下一个
  子树，然后向上导航父节点。在上图中:`1 -> 3 -> 2 -> 5 -> 7 -> 6 -> 4 -> 9 -> 11 -> 10 -> 8`

二叉搜索树由具有属性或属性的节点组成：

- key 整形
- value 整形
- 树的左右节点 leftNode 和 rightNode

一个树节点：

```GO
// TreeNode
type TreeNode struct {
    key int
    value int
    leftNode *TreeNode
    rightNode *TreeNode
}
```

二叉搜索树 BinarySearchTree 由 TreeNode 类型的 rootNode 和 sync.RWMutex 类型的 lock 组成。 通过访问
rootNode 左侧和右侧的节点，可以从 rootNode 遍历二叉搜索树：

```GO
// BinarySearchTree
type BinarySearchTree struct {
  // 根节点
  rootNode *TreeNode
  lock sync.RWMutex
}
```

## Insert 方法

```GO
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
