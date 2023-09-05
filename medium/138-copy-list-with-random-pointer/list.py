from typing import Optional

# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random

class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        c = head

        if c is None:
            return None

        def copyNode(node):
            return Node(node.val)
        
        c.id = 0
        c = c.next
        h = copyNode(head)
        copyHead = h
        copies = [h]
        count = 1
        while c is not None:
            incoming = copyNode(c)
            copies.append(incoming)
            c.id = count
            h.next = incoming
            h = h.next
            c = c.next
            count += 1

        c = head
        h = copyHead

        while c is not None:
            if c.random is not None:
                index = c.random.id
                h.random = copies[index]
            h = h.next
            c = c.next

        return copyHead
