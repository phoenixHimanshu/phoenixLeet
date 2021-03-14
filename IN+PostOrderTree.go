package main

//Problem: Given inorder and pre order tree create a tree
//In-Order Array [3,12,6,4,7,10,11,5,2,8]     Left - Root  - Right
//Post-Order Array [3,6,7,4,12,11,8,2,5,10]   Left - Right - Root

type BNode struct {
	key   int
	left  *BNode
	right *BNode
}

type BTree struct {
	root *Node
}

func buildTree(InOrd []int, PosOrd []int) *BNode {

	len := len(InOrd)

	return helper(InOrd, PosOrd, 0, len-1, 0, len-1)
}

func helper(InOrd []int, PosOrd []int, inStart int, inEnd int, poStart int, poEnd int) *BNode {

	if inStart > inEnd {
		return nil
	}

	rootVal := PosOrd[poEnd]
	root := BNode{rootVal, nil, nil}

	rootIndex := inStart
	for ; InOrd[rootIndex] == rootVal; inStart++ {
		break
	}

	leftSizeTree := rootIndex - inStart
	rightSizeTree := inEnd - rootIndex

	root.left = helper(InOrd, PosOrd, inStart, rootIndex-1, poStart, poStart+leftSizeTree-1)
	root.right = helper(InOrd, PosOrd, rootIndex+1, inEnd, poEnd-rightSizeTree, poEnd-1)

	return &root
}

func main() {

	inOrder := []int{4, 2, 5, 1, 6, 3, 7}
	postOrder := []int{4, 5, 2, 6, 7, 3, 1}

	buildTree(inOrder, postOrder)

}
