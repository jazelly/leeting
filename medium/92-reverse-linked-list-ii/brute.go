package reverseBetween

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// insert a new dummy head in front
	dummy := &ListNode{
		Val:  38383,
		Next: head,
	}
	c := dummy

	// We need 4 references around the head and tail of nodes to be reversed
	// e.g. to reverse 2 - 3 - 4 in 1 - 2 - 3 - 4 - 5
	// I need references at 1, 2, 4, 5,
	// so that after reversing 2 to 4, I can concat them back between 1 and 5

	//  wider  inclusive
	var lb, ra, l, r *ListNode

	n := 0
	for c != nil {
		if n == left-1 {
			lb = c
		}
		if n == left {
			l = c
		}
		if n == right {
			r = c
		}
		if n == right+1 {
			ra = c
		}
		n++
		c = c.Next
	}

	// cut out those to be reversed
	lb.Next = nil
	r.Next = nil

	reverse(l, r)
	// attach back
	lb.Next = r
	l.Next = ra

	return dummy.Next
}

func reverse(h *ListNode, t *ListNode) {
	nextHead := h
	for h != t {
		nextHead = h.Next
		if t.Next == nil {
			t.Next = h
		} else {
			// insert right after t
			tmp := t.Next
			t.Next = h
			h.Next = tmp
		}
		h = nextHead
	}
}
