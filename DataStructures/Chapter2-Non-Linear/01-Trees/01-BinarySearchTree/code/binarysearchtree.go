// Package main provides ...
package binarysearchtree

import (
	"sync"
)

type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

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

func postOrderTraverseTree(treeNode *TreeNode, f func(int)) {
	if treeNode != nil {
		postOrderTraverseTree(treeNode.leftNode, f)
		postOrderTraverseTree(treeNode.rightNode, f)
	}
}

// 后续遍历: 左子树-右子树-根节点
func (bst *BinarySearchTree) PostOrderTraverseTree(treeNode *TreeNode, f func(int)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	postOrderTraverseTree(bst.rootNode, f)
}

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
