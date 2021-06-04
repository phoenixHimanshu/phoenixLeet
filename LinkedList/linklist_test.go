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
	l := ll{nil, 0}
	l1 := ll{nil, 0}
	l2 := ll{nil, 0}

	l1.insertNode(4)
	l1.insertNode(20)

	l2.insertNode(10)
	l2.insertNode(15)
	l2.insertNode(50)
	l2.insertNode(60)
	ll := l.MergeLL(l1, l2)
	for ll.head != nil {
		println(ll.head.data)
		ll.head = ll.head.next
	}

}

func TestMeargeSortedLL1(t *testing.T) {
	l := ll{nil, 0}
	l1 := ll{nil, 0}
	l2 := ll{nil, 0}

	l1.insertNode(4)
	l1.insertNode(20)
	l1.insertNode(50)
	l1.insertNode(60)

	l2.insertNode(10)
	l2.insertNode(15)

	ll := l.MergeLL(l1, l2)
	for ll.head != nil {
		println(ll.head.data)
		ll.head = ll.head.next
	}

}

func TestMeargeSortedLL2(t *testing.T) {
	l := ll{nil, 0}
	l1 := ll{nil, 0}
	l2 := ll{nil, 0}

	l1.insertNode(1)
	l1.insertNode(2)

	l2.insertNode(3)
	l2.insertNode(4)

	ll := l.MergeLL(l1, l2)
	for ll.head != nil {
		println(ll.head.data)
		ll.head = ll.head.next
	}

}
