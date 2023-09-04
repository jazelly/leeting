'''
https://leetcode.com/problems/count-of-interesting-subarrays/solutions/3994883/c-simple-solution-prefix-sum-map/
'''

from collections import defaultdict
from typing import List


class Solution:
    def countInterestingSubarrays(self, nums: List[int], modulo: int, k: int) -> int:
        '''
        Distributive Property: Modulo follows the distributive property
        with respect to addition and multiplication.
        That is, (a + b) % n is congruent to ((a % n) + (b % n)) % n,
        and (a - b) % n is congruent to ((a % n) - (b % n)) % n,
        and (a * b) % n is congruent to ((a % n) * (b % n)) % n
        '''

        # k < M => k % M == k
        # x = sum[0, l-1]
        # y = sum[0, r]
        # y-x = sum[l, r]
        # if   (y-x) % M == k
        # then (y-k) % M == x % M

        n = len(nums)
        ans = 0
        acc = 0
        cnt = defaultdict(int)
        cnt[0] = 1
        # For Example [11, 12, 21, 31], M = 10, K = 1
        # [1, 0, 1, 1] satisfy is 1, otherwise 0
        # [1, 1, 2, 3] accumulated count for [0, r]
        # [1, 1, 2, 3] mod by M
        # target K is 1, 
        for a in nums:
            toa = 1 if a % modulo == k else 0
            acc += toa

            # the property, (y-k) % M == x % M
            ans += cnt[(acc - k) % modulo]
            # cache the y % M, which will be used as x % M
            cnt[acc % modulo] += 1

        return ans
    def countInterestingSubarrays(self, nums: List[int], modulo: int, k: int) -> int:
        '''
        Distributive Property: Modulo follows the distributive property
        with respect to addition and multiplication.
        That is, (a + b) % n is congruent to ((a % n) + (b % n)) % n,
        and (a - b) % n is congruent to ((a % n) - (b % n)) % n,
        and (a * b) % n is congruent to ((a % n) * (b % n)) % n
        '''

        # k < M => k % M == k
        # x = sum[0, l-1]
        # y = sum[0, r]
        # y-x = sum[l, r]
        # (y-x) % M = ((y % M) - (x % M)) % M

        n = len(nums)
        ans = 0
        acc = 0
        cnt = defaultdict(int)
        cnt[0] = 1
        # For Example [11, 12, 21, 31], M = 10, K = 1
        # [1, 0, 1, 1] satisfy is 1, otherwise 0
        # [1, 1, 2, 3] accumulated count for [0, r]
        # [1, 1, 2, 3] mod by M
        # target K is 1, 
        for a in nums:
            toa = 1 if a % modulo == k else 0
            acc += toa
            acc %= modulo

            # distributive property
            acck = acc - k
            acckm = acck % modulo
            acckmc = cnt[acckm]
            ans += acckmc

            cnt[acc] += 1

        print(cnt)

        return ans
    
s = Solution()
nums = [8, 8, 7, 8, 8, 7, 7, 8, 8, 8, 8, 8, 7]
modulo = 2
k = 0
s.countInterestingSubarrays(nums, modulo, k)