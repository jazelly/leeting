class Solution:
    def numTrees(self, n: int) -> int:
        dp = [0, 1, 2]

        for i in range(3, n+1):
            # 1 as root: all right
            r = dp[i-1]

            # 2...i-1 as root: 
            for j in range(2, i):
                r += dp[j-1] * dp[i-j]

            # i as root: all left
            r += dp [i-1]
            dp.append(r)

        return dp[n]