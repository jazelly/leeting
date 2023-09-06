from typing import List
from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def splitListToParts(self, head: Optional[ListNode], k: int) -> List[Optional[ListNode]]:
        n = 0
        c = head
        while c is not None:
            n += 1
            c = c.next
        
        indices = []
        count = n
        kk = k
        while count > 0:
            t = count / kk
            tint = math.ceil(t)
            count -= tint
            kk -= 1
            indices.append(tint)

        while kk > 0:
            indices.append(0)
            kk -= 1
        
        res = []
        c = head
        while len(indices) > 0:
            howMany = indices[0]
            if howMany == 0:
                res.append(None)
            else:
                initial = True
                h = None
                for _ in range(howMany):
                    if initial:
                        h = c
                        initial = False
                    else:
                        c = c.next
                    
                cc = c 
                c = c.next
                cc.next = None
                res.append(h)


            indices = indices[1:]
        return res
