package main

//Implementation of Floyd’s Cycle-Finding Algorithm:
/*
Complexity Analysis:

Time complexity: O(n).
Only one traversal of the loop is needed.
Auxiliary Space:O(1).
There is no space required.

Approach: This is the fastest method and has been described below:

    Traverse linked list using two pointers.
    Move one pointer(slow_p) by one and another pointer(fast_p) by two.
    If these pointers meet at the same node then there is a loop.
	If pointers do not meet then linked list doesn’t have a loop

*/

func (l *ll) DetectCycle() bool {

	slowPtr := l.head
	fastPtr := l.head

	for slowPtr != nil && fastPtr != nil && fastPtr.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			return true
		}
	}

	return false
}

/*
How does above algorithm work?
Please See : How does Floyd’s slow and fast pointers approach work?
https://www.youtube.com/watch?v=Aup0kOWoMVg

References:
http://en.wikipedia.org/wiki/Cycle_detection
http://ostermiller.org/find_loop_singly_linked_list.html
*/
