from typing import List


class Solution:
    def findLongestChain(self, pairs: List[List[int]]) -> int:
        sp = sorted(pairs, key=lambda x: x[1])
        i = 1
        p = 0
        c = 1
        while i< len(sp):
            if sp[i][0] > sp[p][1]:
                c += 1
                p = i
            i += 1

        return c