class Solution:
    def countBits(self, n: int) -> List[int]:
        dp = [0, 1, 1]
        c = 1
        if n == 0:
            return [0]
        if n == 1:
            return [0, 1]

        for i in range(3, n+1):
            if c == len(dp) - c:
                dp.append(1)
                c = 1
            else:
                dp.append(dp[c] + 1)
                c += 1

        return dp
