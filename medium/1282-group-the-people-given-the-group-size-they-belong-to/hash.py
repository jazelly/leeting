from collections import defaultdict
from typing import List


class Solution:
    def groupThePeople(self, groupSizes: List[int]) -> List[List[int]]:
        n = len(groupSizes)
        filled = [False] * n
        fc = 0
        i = 1
        res = []

        group = defaultdict(list)
        for index, v in enumerate(groupSizes):
            group[v].append(index)

        for k, v in group.items():
            if len(v) == k:
                res.append(v)

            elif len(v) > k:
                s = 0
                e = k
                while e <= len(v):
                    res.append(v[s:e])
                    s += k
                    e += k

        return res