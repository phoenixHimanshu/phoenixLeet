package main

import (
	"container/list"
	"strconv"
	"strings"
)

// Leetcode 297

func serialize(root *Node) string {
	if root == nil {
		return "x"
	}

	leftSerialize := serialize(root.left)
	RightSerialize := serialize(root.right)

	return strconv.Itoa(root.key) + "," + leftSerialize + "," + RightSerialize
	//1,2,X,X,3,X,X
}

func deserialize(str string) *Node {

	split := strings.Split(str, ",")
	queue := list.New()

	for _, val := range split {
		queue.PushBack(val)
	}
	root := deserial(queue)

	return root
}

func deserial(ll *list.List) *Node {

	if ll.Len() == 0 {
		return nil
	}

	peek := ll.Front()
	if peek.Value == "x" {
		return nil
	}
	str := peek.Value.(string)
	i, _ := strconv.Atoi(str)
	root := Node{key: i}

	ll.Remove(peek)
	root.left = deserial(ll)
	root.right = deserial(ll)

	return &root
}
