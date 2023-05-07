# this beats 90% of time and 10% of space

class Solution:
    def maxSubArray(self, nums):
        dp = []
        start = -10000 # should be negative INT_MAX
        max_sum = -10000 # should be negative INT_MAX
        the_sum = 0
        for i in range(len(nums)):
            the_sum += nums[i]
            if the_sum >= max_sum:
                max_sum = the_sum
                start = nums[i]
                dp.append(the_sum)
            elif nums[i] > the_sum:
                start = nums[i]
                the_sum = start
                dp.append(the_sum)

        return max(dp)