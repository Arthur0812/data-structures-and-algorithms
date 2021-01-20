package quadtrees

import (
	"fmt"
	"io"
)

// QuadTree represents a quad-tree data structure
type QuadTree struct {
	value          int
	lu, ll, ru, rl *QuadTree
}

// CONSTRUCTORS

// BaseQT returns a single quad-tree with the given label value
func BaseQT(value int) *QuadTree {
	return &QuadTree{value: value}
}

// MakeQT builds a quad-tree from four quad-trees luqt, ruqt, llqt, rlqt
func MakeQT(luqt, ruqt, llqt, rlqt *QuadTree) *QuadTree {
	return &QuadTree{
		lu: luqt,
		ru: ruqt,
		ll: llqt,
		rl: rlqt,
	}
}

// SELECTORS

func value(qt *QuadTree) int {
	return qt.value
}

func lu(qt *QuadTree) *QuadTree {
	return qt.lu
}

func ru(qt *QuadTree) *QuadTree {
	return qt.ru
}

func ll(qt *QuadTree) *QuadTree {
	return qt.ll
}

func rl(qt *QuadTree) *QuadTree {
	return qt.rl
}

// CONDITIONS

func isValue(qt *QuadTree) bool {
	return qt.value != 0
}

// ALGORITHMS

func rotate(qt *QuadTree) *QuadTree {
	if isValue(qt) {
		return qt
	}
	return MakeQT(rotate(lu(qt)), rotate(ru(qt)), rotate(ll(qt)), rotate(rl(qt)))
}

func print(qt *QuadTree, w io.Writer) {
	if isValue(qt) {
		fmt.Fprintf(w, "%d\n", value(qt))
		return
	}
	print(lu(qt), w)
	print(ru(qt), w)
	print(ll(qt), w)
	print(rl(qt), w)
}
