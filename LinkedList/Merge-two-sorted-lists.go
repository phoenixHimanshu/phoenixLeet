package main

func (l *ll) MergeLL(l1 ll, l2 ll) *ll {
	resultll := ll{}
	ll1ptr := l1.head
	ll2ptr := l2.head

	for ll1ptr != nil && ll2ptr != nil {

		if ll1ptr.data == ll2ptr.data {

			resultll.insertNodeSorted(ll1ptr.data, &resultll)

		} else if ll1ptr.data > ll2ptr.data {

			resultll.insertNodeSorted(ll2ptr.data, &resultll)
			resultll.insertNodeSorted(ll1ptr.data, &resultll)

		} else {
			resultll.insertNodeSorted(ll1ptr.data, &resultll)
			resultll.insertNodeSorted(ll2ptr.data, &resultll)

		}

		ll1ptr = ll1ptr.next
		ll2ptr = ll2ptr.next
	}

	diff := l1.len - l2.len
	var biggerll ll
	var travlen int
	if diff > 0 {
		biggerll = l1
		travlen = l2.len
	} else {
		biggerll = l2
		diff = l2.len - l1.len
		travlen = l1.len
	}
	runner := biggerll.head
	for runner != nil {

		if travlen != 0 {

			travlen--
			runner = runner.next
			continue
		}

		resultll.insertNode(runner.data)
		runner = runner.next

	}

	return &resultll
}

func (l *ll) insertNodeSorted(val int, resultll *ll) {
	n := node{}
	n.data = val
	if resultll.len == 0 {
		resultll.head = &n
		resultll.len++
		return
	}

	ptr := resultll.head
	var prev *node
	for i := resultll.len; i != 0; i-- {

		if ptr.data > val {

			//New node becomes the first node
			if prev == nil {
				resultll.head = &n
				resultll.len++
				resultll.head.next = ptr
				break
			} else if prev.data < val {

				//Middle node case
				temp := ptr
				ptr = &n
				prev.next = ptr
				ptr.next = temp
				resultll.len++
				break

			} else if ptr.next == nil {
				//last node case
				ptr.next = &n
				resultll.len++
				break
			}

		} else if ptr.data < val && resultll.len != 1 {
			prev = ptr
			//Add to the End of this is the last val to be inserted
			if i == 1 {
				ptr.next = &n
				n.next = nil
				resultll.len++
			}
			ptr = ptr.next
			continue

		} else {
			//Last node addition
			ptr.next = &n
			n.next = nil
			resultll.len++
		}

		prev = ptr
		ptr = ptr.next

	}

}
