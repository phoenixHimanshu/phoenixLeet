package main

import (
	"fmt"
	"testing"
)

func TestInPostOrderTree(t *testing.T) {

	inOrder := []int{4, 2, 5, 1, 6, 3, 7}
	postOrder := []int{4, 5, 2, 6, 7, 3, 1}

	root := buildTree(inOrder, postOrder)

	InOrderTraversal(root)
	//t.Errorf("Test case failed")

	//t.Error()
}

func TestSerializeTree(t *testing.T) {

	var tree Tree

	tree.insert(2)
	tree.insert(1)
	tree.insert(3)

	strStream := serialize(tree.root)
	fmt.Println(strStream)
	InOrderTraversal(deserialize(strStream))
	//Todo not working 1,2 is working fine but 3 is missing
	// solve the bug

}

func TestNaryTree(t *testing.T) {
	// 6 node tree
	root := TreeNode{2, nil}
	node1 := TreeNode{3, nil}
	node2 := TreeNode{4, nil}
	node3 := TreeNode{5, nil}
	node4 := TreeNode{6, nil}
	node5 := TreeNode{7, nil}
	node6 := TreeNode{8, nil}
	//node3 := TreeNode{6,nil}

	// 5 edges required to build 6 node tree
	root.children = append(root.children, &node1)
	root.children = append(root.children, &node2)
	root.children = append(root.children, &node3)
	node2.children = append(node2.children, &node4)
	node2.children = append(node2.children, &node5)
	node4.children = append(node4.children, &node6)

	//i:=depthOfTree(&root)
	//fmt.Println(i)

	j := maxDepth(&root)
	fmt.Println(j)

	printTree(root)

}

func TestBinaryTreeTraversal(t *testing.T) {

	var tree Tree
	var p Tree
	tree.insert(2)
	tree.insert(3)
	tree.insert(1)
	p.insert(2)
	p.insert(3)
	p.insert(1)

	//fmt.Println(isSameTree(t.root,p.root))
	fmt.Println("PreOrderTraversal")
	preOrderTraversal(tree.root)
	fmt.Println("PostOrderTraversal")
	postOrderTraversal(tree.root)
	fmt.Println("InOrderTraversal")
	InOrderTraversal(tree.root)

}
