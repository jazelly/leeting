package add2Numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var c *ListNode
	var head *ListNode

	buffer := 0
	for l1 != nil && l2 != nil {
		r := l1.Val + l2.Val
		v2a := r % 10

		if buffer > 0 {
			v2a += buffer
			buffer = 0
			if v2a/10 > 0 {
				buffer += v2a / 10
				v2a = v2a % 10
			}
		}
		theNew := &ListNode{
			Val: v2a,
		}

		// first one
		if c == nil || head == nil {
			head = theNew
			c = theNew
		} else {
			c.Next = theNew
			c = c.Next
		}

		if b := r / 10; b > 0 {
			buffer += b
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		if buffer > 0 {
			l1.Val += buffer
			buffer = 0
			if l1.Val/10 > 0 {
				buffer += l1.Val / 10
				l1.Val = l1.Val % 10
			}
		}
		c.Next = l1
		c = c.Next
		l1 = l1.Next
	}

	for l2 != nil {
		if buffer > 0 {
			l2.Val += buffer
			buffer = 0
			if l2.Val/10 > 0 {
				buffer += l2.Val / 10
				l2.Val %= 10
			}
		}
		c.Next = l2
		c = c.Next
		l2 = l2.Next
	}

	if buffer > 0 {
		last := &ListNode{
			Val: buffer,
		}
		c.Next = last
	}

	return head
}
