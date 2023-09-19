from collections import defaultdict
from typing import List


class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        mp = defaultdict(bool)
        for n in nums:
            if mp[n]:
                return n
            mp[n] = True
        return 0
