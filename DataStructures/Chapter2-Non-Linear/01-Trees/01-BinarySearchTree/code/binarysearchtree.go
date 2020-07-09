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

func (bst *BinarySearchTree) PostOrderTraverseTree(treeNode *TreeNode, f func(int)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	postOrderTraverseTree(bst.rootNode, f)
}

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
