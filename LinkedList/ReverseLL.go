package main

import "fmt"

type node struct {
	data int
	next *node
}

type ll struct {
	head *node
	len  int
}

func (l *ll) insertNode(val int) {
	n := node{}
	n.data = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}

	ptr := l.head
	for i := l.len; i != 0; i-- {

		if ptr.next == nil {
			ptr.next = &n
			n.next = nil
			l.len++
		}

		ptr = ptr.next
	}

}

func (l *ll) printll() {
	ptr := l.head
	for i := l.len; i != 0; i-- {
		fmt.Println(ptr.data)

		ptr = ptr.next
	}
}

func (l *ll) reverse() {
	runner := l.head
	var prev *node
	for i := 0; i < l.len; i++ {
		next := runner.next
		runner.next = prev
		prev = runner
		runner = next

	}
	l.head = prev

}

func main() {
	l := ll{nil, 0}
	l.insertNode(10)
	l.insertNode(20)
	l.insertNode(30)
	l.reverse()
	l.printll()

}
