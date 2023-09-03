'''
https://leetcode.com/contest/weekly-contest-361/problems/count-of-interesting-subarrays/
'''
# This will TLE for large input
class Solution:
    def countInterestingSubarrays(self, nums: List[int], modulo: int, k: int) -> int:
        n = len(nums)
        count = 0
        
        # make a leading 0 table
        cc = 0
        tc = 0
        leadingl = []
        for i, v in enumerate(nums):
            if v % modulo == k:
                tc += 1
                leadingl.append((i, cc))
                cc = 0
                continue
            cc += 1
        leadingl.append((n, cc))                

        nums10 = [0] * n
        for i, v in enumerate(nums):
            if v % modulo == k:
                nums10[i] = 1


        # [0, 0, 1, 0, 0, 0, 1, 0, 0, 1]
        m = len(leadingl)
        # [(2, 2), (6, 3), (9, 2), (10, 0)]
        count = 0
        b = 0
        while modulo * b + k <= tc:
            target = modulo * b + k
            if target > 0:
                i = 0
                while i + target < m:
                    # prev 0
                    p0 = leadingl[i][1] + 1
                    # tail 0
                    t0 = leadingl[i+target][1] + 1
                    count += p0 * t0
                    i += 1
            b += 1
    
        if k == 0:
            for l in leadingl:
                count += (l[1] * (l[1] + 1))//2
      
        return count
    
s = Solution()
nums = [3,1,9,6]
modulo = 3
k = 0
s.countInterestingSubarrays(nums, modulo, k)