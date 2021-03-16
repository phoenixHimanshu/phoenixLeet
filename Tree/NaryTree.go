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
