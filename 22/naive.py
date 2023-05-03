# beats 5% on space and time

from typing import List

class Node:
    def __init__(self):
        self.parent = None
        self.left = None
        self.right = None
        self.value = "("
        self.previousLeft = 1  # inclusive
        self.previousRight = 0  # inclusive
        self.height = 1

    def trace_back(self):
        l = []
        current = self
        while current.parent != None:
            l.append(current.value)
            current = current.parent
        l.append(current.value)
        return "".join(reversed(l))


class Solution:
    def generateNode(self, l, n):
        cl = []
        for p in l:
            no_right_child = False

            # closure
            if p.previousLeft > 0 and p.previousRight == p.previousLeft:
                no_right_child = True

            # generate a left child
            h = p.height + 1
            if p.previousLeft < n:
                child1 = Node()
                child1.height = h
                child1.previousLeft = p.previousLeft + 1
                child1.previousRight = p.previousRight
                child1.parent = p
                p.left = child1
                if h != max:
                    cl.append(child1)

            # generate a right child
            if not no_right_child and p.previousRight < n:
                child2 = Node()
                child2.value = ")"
                child2.height = h
                child2.previousLeft = p.previousLeft
                child2.previousRight = p.previousRight + 1
                child2.parent = p
                p.right = child2
                cl.append(child2)
        return cl

    def generateParenthesis(self, n: int) -> List[str]:
        root = Node()
        l = [root]
        for i in range(n * 2 - 1):
            l = self.generateNode(l, n)
        result = [item.trace_back() for item in l]
        return result
