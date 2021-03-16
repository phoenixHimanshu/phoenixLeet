package main

import "testing"

func TestReverseLL(t *testing.T) {
	l := ll{nil, 0}
	l.insertNode(10)
	l.insertNode(20)
	l.insertNode(30)
	l.reverse()
	l.printll()

}
