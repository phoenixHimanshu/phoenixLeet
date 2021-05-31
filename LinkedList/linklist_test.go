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

func TestDetectCycle(t *testing.T) {
	l := ll{nil, 0}
	l.insertNode(20)
	l.insertNode(4)
	l.insertNode(15)
	l.insertNode(10)
	//create a cycle in linklist
	l.head.next.next.next.next = l.head
	println(l.DetectCycle())
}

func TestDetectCycle1(t *testing.T) {
	l := ll{nil, 0}
	l.insertNode(20)
	l.insertNode(4)
	l.insertNode(15)
	l.insertNode(10)
	println(l.DetectCycle())
}

func TestMeargeSortedLL(t *testing.T) {
	l1 := ll{nil, 0}
	l2 := ll{nil, 0}

	l1.insertNode(20)
	l1.insertNode(4)

	l2.insertNode(15)
	l2.insertNode(10)
	ll := MergeLL(l1, l2)
	for ll.next != nil {
		println(ll.data)
	}
}
