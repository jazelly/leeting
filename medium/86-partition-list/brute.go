package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var lc, rc, lhead, rhead *ListNode

	cursor := head
	for cursor != nil {
		theNext := cursor.Next
		// cursor.Next = nil

		if cursor.Val < x {
			if lc == nil {
				lc = cursor
				lhead = cursor
			} else {
				lc.Next = cursor
				lc = lc.Next
			}
		} else {
			if rc == nil {
				rhead = cursor
				rc = cursor
			} else {
				rc.Next = cursor
				rc = rc.Next
			}
		}

		cursor = theNext
	}

	if lc == nil {
		return rhead
	}

	lc.Next = rhead
	return lhead
}

func main() {
	head := &ListNode{Val: 2}
	one := &ListNode{Val: 1}
	head.Next = one
	h := partition(head, 2)
	for h != nil {
		fmt.Println(h.Val)
	}
}
