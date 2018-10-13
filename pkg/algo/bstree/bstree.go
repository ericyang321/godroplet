package bstree

import (
	"fmt"
	"sync"
)

// BSTree is a concurrency safe, iteratable binary search tree
type BSTree struct {
	lock sync.Mutex
	root *node
}

type node struct {
	key   string
	val   interface{}
	left  *node
	right *node
}

// Node is a node in binary search tree
type Node struct {
	Key string
	Val interface{}
}

// Insert inserts new information by creating a node or updating
// an existing node
func (b *BSTree) Insert(key string, val interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.root == nil {
		b.root = &node{key: key, val: val}
		return
	}
	b.root.insert(key, val)
}

// Iterate creates a read-only channel that outputs nodes in order
func (b *BSTree) Iterate() <-chan Node {
	ch := make(chan Node)
	b.lock.Lock()
	go func() {
		b.root.iterate(ch)
		b.lock.Unlock()
		close(ch)
	}()
	return ch
}

func (n *node) insert(key string, val interface{}) {
	// traverse
	if key < n.key {
		if n.left == nil {
			n.left = &node{key: key, val: val}
		} else {
			n.left.insert(key, val)
		}
	} else if key > n.key {
		if n.right == nil {
			n.right = &node{key: key, val: val}
		} else {
			n.right.insert(key, val)
		}
	} else {
		// act
		n.val = val
	}
}

// value recursively iterates nodes, depth first, until a node is found
func (n *node) value(key string) (interface{}, error) {
	if n == nil {
		return nil, fmt.Errorf("Node of key %s doesn't exist", key)
	}
	// traverse
	if key < n.key {
		return n.left.value(key)
	} else if key > n.key {
		return n.right.value(key)
	}
	// act
	return n.value, nil
}

func (n *node) iterate(ch chan<- Node) {
	if n == nil {
		return
	}
	n.left.iterate(ch)
	ch <- Node{
		Key: n.key,
		Val: n.val,
	}
	n.right.iterate(ch)
}

func (n *node) remove(key string) (*node, error) {
	var err error
	if n == nil {
		return nil, fmt.Errorf("Node of key %s doesn't exist", key)
	}
	// traverse
	if key < n.key {
		n.left, err = n.left.remove(key)
		return n, err
	}
	if key > n.key {
		n.right, err = n.right.remove(key)
		return n, err
	}
	// act
	// node is a leaf node
	if n.isLeaf() {
		return nil, nil
	}
	// node still has a left child left
	if n.hasLeft() && !n.hasRight() {
		return n.left, nil
	}
	// node still has a right child left
	if !n.hasLeft() && n.hasRight() {
		return n.right, nil
	}
	// node has right and left children left
	minNode := n.right.minNode()
	n.key = minNode.key
	n.val = minNode.val
	n.right, err = n.right.remove(minNode.key)
	return n, err
}

func (n *node) isLeaf() bool {
	return !n.hasLeft() && !n.hasRight()
}

func (n *node) hasLeft() bool {
	return n.left != nil
}

func (n *node) hasRight() bool {
	return n.right != nil
}

func (n *node) minNode() *node {
	for n.left != nil {
		n = n.left
	}
	return n
}
