class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        dp = [[0 for _ in range(n)] for _ in range(m)]

        dp[0][0] = 1
        i = 0
        while i < m:
            j = 0
            while j < n:
                if i - 1 >= 0:
                    dp[i][j] += dp[i-1][j]
                if j - 1 >= 0:
                    dp[i][j] += dp[i][j-1]
                j += 1
            i += 1

        return dp[m-1][n-1]