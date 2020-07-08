## Tree 树

树是一种非线性数据结构。 树用于搜索和其他用例。 二叉树的节点最多有两个子节点。 二叉搜索树由左侧节点的
属性值小于右侧节点的属性值的节点组成。 根节点位于树的零级。 每个子节点可以是叶。没有节点的树被称为空树

## Binary search tree 二叉搜索树

二叉搜索树，也称为二叉查找树、有序二叉树（ordered binary tree）或排序二叉树（sorted binary tree），
是指一棵空树或者具有下列性质的二叉树：

- 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
- 若任意节点的右子树不空，则右子树上所有节点的值均大于或等于它的根节点的值；
- 任意节点的左、右子树也分别为二叉搜索树；

如下图，左侧的就是一棵二叉搜索树;而右侧的则不是，因为节点 9 在根节点 10 的右子树上，但
是其值却比 10 小.

![BinarySearchTree](/image/binarysearchtree.png)

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

## InsertElement 插入元素方法

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
func (bst *BinarySearchTree) InsertElement(key, value int) {
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
