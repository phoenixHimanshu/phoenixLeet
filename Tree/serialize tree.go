package main

import "strconv"

func serialize(root *Node) string {
	if root == nil {
		return "x"
	}

	leftSerialize := serialize(root.left)
	RightSerialize := serialize(root.right)

	return strconv.Itoa(root.key) + "," + leftSerialize + "," + RightSerialize
}
