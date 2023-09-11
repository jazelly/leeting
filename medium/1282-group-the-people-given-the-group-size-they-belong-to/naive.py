from typing import List


class Solution:
    def groupThePeople(self, groupSizes: List[int]) -> List[List[int]]:
        n = len(groupSizes)
        filled = [False] * n
        fc = 0
        i = 1
        res = []

        while fc < n:
            buffer = []
            redo = False
            for index, v in enumerate(groupSizes):
                if not filled[index] and v == i and len(buffer) < i:
                    buffer.append(index)
                    filled[index] = True
                    fc += 1

                elif not filled[index] and v == i:
                    redo = True

            if len(buffer) > 0:
                res.append(buffer)

            if not redo:
                i += 1

        return res