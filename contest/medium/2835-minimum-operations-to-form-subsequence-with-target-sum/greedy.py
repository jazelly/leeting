from typing import List

# https://leetcode.com/contest/weekly-contest-360/problems/minimum-operations-to-form-subsequence-with-target-sum/
class Solution:
    def minOperations(self, nums: List[int], target: int) -> int:
        total = sum(nums)
        if total < target:
            return -1
        nums.sort()
        o = 0
        while target:
            num = nums.pop()
            if total - num >= target:
                total -= num
            elif num > target:
                o += 1
                nums.append(num//2)
                nums.append(num//2)
            else:
                target -= num
                total -= num
        return o
    