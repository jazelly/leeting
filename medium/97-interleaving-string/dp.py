class Solution:
    def isInterleave(self, s1: str, s2: str, s3: str) -> bool:
        m = len(s1)
        n = len(s2)

        if m+n != len(s3):
            return False

        if len(s3) == 0:
            return True

        dp = [[-1 for _ in range(n+1)] for _ in range(m+1)]
        dp[0][0] = 1
        v = [[False for _ in range(n+1)] for _ in range(m+1)]
        q = [(0, 0)]
        v[0][0] = True

        while len(q) > 0:
            node = q[0]
            x = node[0]
            y = node[1]

            if dp[x][y] < 0:
                target = s3[x+y-1]
                if x - 1 >= 0 and dp[x-1][y] == 1:
                    if s1[x-1] == target:
                        dp[x][y] = 1
                
                if y - 1 >= 0 and dp[x][y-1] == 1:
                    if s2[y-1] == target:
                        dp[x][y] = 1

                if dp[x][y] != 1:
                    dp[x][y] = 0
                v[x][y] = True
            if x + 1 <= m and not v[x+1][y]:
                v[x+1][y] = True
                q.append((x+1, y))
            if y + 1 <= n and not v[x][y+1]:
                v[x][y+1] = True
                q.append((x, y+1))

            q = q[1:]

        return dp[m][n] == 1