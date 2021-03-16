package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

// Tree
func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

// Node
func (n *Node) insert(data int) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func isSameTree(p *Node, q *Node) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.key != q.key {
		return false
	}
	return isSameTree(p.left, q.left) && isSameTree(p.right, q.right)
}

func preOrderTraversal(root *Node) {

	if root == nil {
		return
	}
	fmt.Println(root.key)
	preOrderTraversal(root.left)
	preOrderTraversal(root.right)

}

func postOrderTraversal(root *Node) {

	if root == nil {
		return
	}
	preOrderTraversal(root.left)
	preOrderTraversal(root.right)
	fmt.Println(root.key)
}

func InOrderTraversal(root *Node) {

	if root == nil {
		return
	}
	preOrderTraversal(root.left)
	fmt.Println(root.key)
	preOrderTraversal(root.right)

}
