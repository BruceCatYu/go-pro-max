package gopromax

import "cmp"

type Tree[T cmp.Ordered] struct {
	v    T
	root *Node[T]
}

func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(v)
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(v)
}

type Node[T cmp.Ordered] struct {
	val         T
	left, right *Node[T]
}

func (n *Node[T]) Add(v T) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}
	t := cmp.Compare(v, n.val)
	if t < 0 {
		n.left = n.left.Add(v)
	}
	if t > 0 {
		n.right = n.right.Add(v)
	}
	return n
}

func (n *Node[T]) Contains(v T) bool {
	if n == nil {
		return false
	}
	t := cmp.Compare(v, n.val)
	if t < 0 {
		return n.left.Contains(v)
	}
	if t > 0 {
		return n.right.Contains(v)
	}
	return true
}

func NewTree[T cmp.Ordered](v T) *Tree[T] {
	return &Tree[T]{
		v: v,
	}
}
