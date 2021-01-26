package trees

type node struct {
	value       int
	left, right *node
}

// BST represents a binary search tree data structure
type BST struct {
	root *node
}

// EmptyTree is an empty binary tree
var EmptyTree = &BST{}

// MakeTree creates a new tree with the given value as the root node's value, and puts the left and right trees as its subtrees
func MakeTree(value int, left *BST, right *BST) *BST {
	n := &node{
		value: value,
		left:  left.root,
		right: right.root,
	}
	return &BST{root: n}
}

// SELECTORS (panic when used on EmptyTree)

func root(bst *BST) int {
	return bst.root.value
}

func left(bst *BST) *BST {
	return &BST{root: bst.root.left}
}

func right(bst *BST) *BST {
	return &BST{root: bst.root.right}
}

// CONDITIONS

func isEmpty(bst *BST) bool {
	return bst.root == nil
}

// ALGORITHMS

func insert(value int, bst *BST) *BST {
	if isEmpty(bst) { // base case
		return MakeTree(value, EmptyTree, EmptyTree)
	}
	if value < root(bst) { // reprocess OP on left subtree
		return MakeTree(root(bst), insert(value, left(bst)), right(bst))
	}
	if value > root(bst) { // reprocess OP on the right subtree
		return MakeTree(root(bst), left(bst), insert(value, right(bst)))
	}
	panic("inserting: error: violated assumption in procedure insert")
}

func isIn(value int, bst *BST) bool {
	for !isEmpty(bst) && value != root(bst) {
		if value < root(bst) {
			bst = left(bst)
		} else {
			bst = right(bst)
		}
	}
	return !isEmpty(bst)
}

func delete(value int, bst *BST) *BST {
	if isEmpty(bst) {
		return nil
	}
	if value < root(bst) { // delete from left subtree
		return MakeTree(root(bst), delete(value, left(bst)), right(bst))
	}
	if value > root(bst) { // delete from right subtree
		return MakeTree(root(bst), left(bst), delete(value, right(bst)))
	}
	// now the root holds the searched value
	if isEmpty(left(bst)) {
		return right(bst)
	}
	if isEmpty(right(bst)) {
		return left(bst)
	}
	// both subtrees are there
	return MakeTree(smallestNode(right(bst)), left(bst), removeSmallestNode(right(bst)))
}

func smallestNode(bst *BST) int {
	if isEmpty(bst) {
		panic("trying to get smallest node of empty tree")
	}
	if isEmpty(left(bst)) {
		return root(bst)
	}
	return smallestNode(left(bst))
}

func removeSmallestNode(bst *BST) *BST {
	if isEmpty(bst) {
		panic("trying to remove smallest node of of empty tree")
	}
	if isEmpty(left(bst)) {
		return right(bst)
	}
	return MakeTree(root(bst), removeSmallestNode(left(bst)), right(bst))
}

func isBST(t *BST) bool {
	if isEmpty(t) { // every empty tree is also a valid binary search tree
		return true
	}
	return allSmaller(left(t), root(t)) && isBST(left(t)) &&
		allBigger(right(t), root(t)) && isBST(right(t))
}

func allSmaller(t *BST, value int) bool {
	if isEmpty(t) {
		return true
	}
	return root(t) < value && allSmaller(left(t), value) && allSmaller(right(t), value)
}

func allBigger(t *BST, value int) bool {
	if isEmpty(t) {
		return true
	}
	return root(t) > value && allBigger(right(t), value) && allBigger(left(t), value)
}
