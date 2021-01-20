package binarytrees

import "testing"

func TestBinaryTree(t *testing.T) {
	testTreeLeft := MakeTree(3, Leaf(1), MakeTree(6, EmptyTree, Leaf(7)))
	testTreeRight := MakeTree(11, MakeTree(9, EmptyTree, Leaf(10)), MakeTree(14, Leaf(12), Leaf(15)))
	testTree := MakeTree(8, testTreeLeft, testTreeRight)

	if root(testTree) != 8 {
		t.Fatalf("binary trees: incorrect value: expected: %v, got: %v", 8, root(testTree))
	}

	if left(testTree).root != testTreeLeft.root {
		t.Fatalf("binary trees: incorrect left tree: expected: %v, got: %v", *testTreeLeft.root, *left(testTree).root)
	}

	if right(testTree).root != testTreeRight.root {
		t.Fatalf("binary trees: incorrect right tree: expected: %v, got: %v", *testTreeRight.root, *right(testTree).root)
	}

	expectedSize := 11
	if s := size(testTree); s != expectedSize {
		t.Fatalf("binary trees: incorrect size determination: expected: %v, got: %v", expectedSize, s)
	}
}
