package main

import "fmt"

type binarySearchTree struct {
	root *node
}

type node struct {
	data   int
	parent *node
	left   *node
	right  *node
}

func (b *binarySearchTree) insertValues(slice []int) {
	for _, value := range slice {
		b.insert(value, b.root)
	}
}

func (b *binarySearchTree) insert(value int, n *node) {
	if value > n.data {
		if n.right != nil {
			b.insert(value, n.right)
		} else {
			n.right = &node{parent: n, data: value}
		}
	} else {
		if n.left != nil {
			b.insert(value, n.left)
		} else {
			n.left = &node{parent: n, data: value}
		}
	}
}

func (b *binarySearchTree) delete(value int, n *node) bool {
	toDelete := b.search(value, n)

	if toDelete == nil {
		return false
	}

	if toDelete.left == nil && toDelete.right == nil {
		return replaceNodeWithChild(toDelete, nil, toDelete.parent)
	}

	if toDelete.left != nil && toDelete.right == nil {
		return replaceNodeWithChild(toDelete, toDelete.left, toDelete.parent)
	} else if toDelete.left == nil && toDelete.right != nil {
		return replaceNodeWithChild(toDelete, toDelete.right, toDelete.parent)
	}

	if toDelete.left != nil && toDelete.right != nil {
		successorNode := findLeftmostNode(toDelete.right)
		if successorNode.right != nil {
			replaceNodeWithChild(successorNode, successorNode.right, successorNode.parent)
		}
		return replaceNodeWithSuccessor(toDelete, successorNode, toDelete.parent)
	}
	return false
}

func findLeftmostNode(n *node) *node {
	if n.left != nil {
		return findLeftmostNode(n.left)
	}
	return n
}

func replaceNodeWithSuccessor(node, successor, parent *node) bool {
	replaceNodeWithChild(node, successor, parent)
	successor.left = node.left
	if node.right != successor {
		successor.right = node.right
	}
	return true
}

func replaceNodeWithChild(node, newChild, parent *node) bool {
	if newChild != nil {
		newChild.parent = parent
	}
	if parent.left == node {
		parent.left = newChild
		return true
	} else if parent.right == node {
		parent.right = newChild
		return true
	}
	return false
}

func (b *binarySearchTree) search(value int, n *node) *node {
	if value < n.data {
		if n.left != nil {
			return b.search(value, n.left)
		}
		return nil
	} else if value > n.data {
		if n.right != nil {
			return b.search(value, n.right)
		}
		return nil
	}

	return n
}

func (b *binarySearchTree) invertNodes(n *node) {
	left := n.left
	right := n.right

	n.left = right
	n.right = left
	if left != nil {
		b.invertNodes(left)
	}
	if right != nil {
		b.invertNodes(right)
	}
}

func (b *binarySearchTree) inOrderTraversal(n *node, callback func(node *node)) {
	if n.left != nil {
		b.inOrderTraversal(n.left, callback)
	}
	callback(n)
	if n.right != nil {
		b.inOrderTraversal(n.right, callback)
	}
}

func printNodeData(n *node) {
	fmt.Println(n.data)
}

func main() {
	myArr := []int{5, 9, 2, 4, 12, 10, 13, 3}

	myTree := binarySearchTree{root: &node{data: 1}}
	myTree.insertValues(myArr)
	myTree.inOrderTraversal(myTree.root, printNodeData)

	myTree.delete(12, myTree.root)

	myTree.invertNodes(myTree.root)
	myTree.inOrderTraversal(myTree.root, printNodeData)
}
