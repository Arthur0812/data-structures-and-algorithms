package binarytrees

type node struct {
	value       int
	left, right *node
}

// BinaryTree represents a binary tree data structure
type BinaryTree struct {
	root *node
}

// CONSTRUCTORS

// EmptyTree is simply an empty binary tree
var EmptyTree = &BinaryTree{}

// MakeTree builds a binary tree from a root node with label "value" and two constituent binary trees "left" and "right"
func MakeTree(value int, left, right *BinaryTree) *BinaryTree {
	newTree := &BinaryTree{
		root: &node{
			value: value,
			left:  left.root,
			right: right.root,
		},
	}
	return newTree
}

// SELECTORS

func root(bt *BinaryTree) int {
	return bt.root.value
}

func left(bt *BinaryTree) *BinaryTree {
	return &BinaryTree{
		root: bt.root.left,
	}
}

func right(bt *BinaryTree) *BinaryTree {
	return &BinaryTree{
		root: bt.root.right,
	}
}

// CONDITIONS

func isEmpty(bt *BinaryTree) bool {
	return bt.root == nil
}

// DERIVED OPERATORS

// Leaf creates a tree only holding only a node with a value and no children
func Leaf(value int) *BinaryTree {
	return MakeTree(value, EmptyTree, EmptyTree)
}

// ALGORITHMS

func size(bt *BinaryTree) int {
	if isEmpty(bt) {
		return 0
	}
	return 1 + size(left(bt)) + size(right(bt))
}
