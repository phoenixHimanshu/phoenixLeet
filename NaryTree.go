package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	data     int
	children []*TreeNode
}

//var maxDepth int64

func printTree(treeNode TreeNode) {
	s := strconv.Itoa(treeNode.data) + " Children ==>"

	for i := 0; i < len(treeNode.children); i++ {
		s = s + strconv.Itoa(treeNode.children[i].data)
	}

	fmt.Println(s)

	for i := 0; i < len(treeNode.children); i++ {

		printTree(*treeNode.children[i])

	}

}

//Avinash Solution
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	maxdepth := 0
	for _, ch := range node.children {
		maxdepth = max(maxDepth(ch), maxdepth)
	}
	return 1 + max(maxdepth, 0)
}

/*
// My Solution
func depthOfTree(treeNode *TreeNode) int64 {
		if treeNode == nil {
			return 0
		}

		helper(treeNode,1)
		return maxDepth
}

func helper(treeNode *TreeNode,level int64)  {

	if treeNode == nil {
		return
	}
	if len(treeNode.children) == 0{
		if maxDepth < level{
			maxDepth = level
		}

	}

	childs := treeNode.children
	for i:=0; i < len(childs); i++{
		helper(childs[i], level+1)
	}

}*/

func main() {

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

	//printTree(root)

}
