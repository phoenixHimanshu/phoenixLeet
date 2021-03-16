package main

//Problem: Given inorder and pre order tree create a tree
//In-Order Array [3,12,6,4,7,10,11,5,2,8]     Left - Root  - Right
//Post-Order Array [3,6,7,4,12,11,8,2,5,10]   Left - Right - Root

func buildTree(InOrd []int, PosOrd []int) *Node {

	length := len(InOrd)

	return helper(InOrd, PosOrd, 0, length-1, 0, length-1)
}

func helper(InOrd []int, PosOrd []int, inStart int, inEnd int, poStart int, poEnd int) *Node {

	if inStart > inEnd {
		return nil
	}

	rootVal := PosOrd[poEnd]
	root := Node{rootVal, nil, nil}

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
