# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        if left == right:
            return head

        dh = ListNode()
        dh.next = head

        c = dh
        i = 0
        while i < left-1:
            c = c.next
            i += 1

        cc = c # the node before cut point
        c = c.next # the cut node included in the reverse range

        i += 1
        reverseTail = c
        reverseHead = c

        c = c.next
        while i < right:
            nrn = c.next

            c.next = reverseHead
            reverseHead = c
            cc.next = reverseHead
            
            c = nrn

            i += 1

         
        reverseTail.next = c
        cc.next = reverseHead
        return dh.next