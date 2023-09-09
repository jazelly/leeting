class Solution:
    def combinationSum4(self, nums: List[int], target: int) -> int:
        dp = [0] * (target+1)
        dp[0] = 1     
        i = 1
        while i <= target:
            for n in nums:
                tl = i - n
                if tl >= 0:
                    dp[i] += dp[tl]
            i += 1

        return dp[target]
