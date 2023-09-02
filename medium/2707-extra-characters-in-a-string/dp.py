from typing import List


class Solution:
    def minExtraChar(self, s: str, dictionary: List[str]) -> int:
        n = len(s)
        dp = [[0 for _ in range(n+1)] for _ in range(n)]

        i = 0
        while i < n:
            for d in dictionary:
                if i + len(d) <= n and d == s[i:i+len(d)] and len(d) > dp[i][i+len(d)]:
                    dp[i][i+len(d)] = len(d)
            i += 1

        for i in range(n):
            for j in range(n+1):
                mx = 0
                for ij in range(i, j):
                    if dp[i][ij] + dp[ij][j] > mx:
                        mx = dp[i][ij] + dp[ij][j]
                if mx > dp[i][j]:
                    dp[i][j] = mx
                
        return n - dp[0][n]
