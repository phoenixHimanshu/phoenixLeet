package main

func (l *ll) MergeLL(l1 ll, l2 ll) node {
	resultll := ll{}
	resultptr := resultll.head
	ll1ptr := l1.head
	ll2ptr := l2.head

	for ll1ptr != nil || ll2ptr != nil {

		for resultll.len > 0 {
			if ll1ptr.data == ll2ptr.data {

				resultll.insertNodeSorted(ll1ptr.data, &resultll)

			} else if ll1ptr.data > ll2ptr.data {

				resultll.insertNodeSorted(ll2ptr.data)
				resultll.insertNodeSorted(ll1ptr.data)

			} else {
				resultll.insertNodeSorted(ll1ptr.data)
				resultll.insertNodeSorted(ll2ptr.data)

			}
		}
		ll1ptr = ll1ptr.next
		ll2ptr = ll2ptr.next
	}

	return *resultptr
}

func insertNodeSorted(val int, resultll *ll) {
	n := node{}
	n.data = val
	if resultll.len == 0 {
		resultll.head = &n
		resultll.len++
		return
	}

	ptr := resultll.head
	for i := resultll.len; i != 0; i-- {

		if ptr.data > val {

		} else {
			ptr.next = &n
			n.next = nil
			resultll.len++
		}

		ptr = ptr.next
	}

}
