from typing import List


class Solution:
    def getMaxFunctionValue(self, receiver: List[int], k: int) -> int:
        res = [i for i in range(len(receiver))]
        
        # find the max index
        def fillRes(mi: int):
            kk = k
            dp = [-1 for _ in range(len(receiver))]
            ci = mi
            c = 0
            cc = 0
            while dp[ci] < 0:
                if kk == 0:
                    res[mi] += cc
                    return
                cn = receiver[ci]
                c += 1
                cc += receiver[ci]
                dp[ci] = receiver[ci]
                ci = cn
                kk -= 1

            if kk == 0:
                res[mi] += cc
                return
            initialCount = c
            initialSum = cc

            c = 0
            cc = 0
            dp = [-1 for _ in range(len(receiver))]
            while dp[ci] < 0:
                cn = receiver[ci]
                c += 1
                cc += receiver[ci]
                dp[ci] = receiver[ci]
                ci = cn

            remain = kk % c
            rounds = kk // c

            ca = ci
            remainCount = 0
            for _ in range(remain):
                remainCount += receiver[ca]
                ca = receiver[ca]
            res[mi] += rounds * cc + remainCount + initialSum
            return
        for i in range(len(receiver)):
            fillRes(i)
        return max(res)
    

        

        

s = Solution()

s.getMaxFunctionValue([0], 2)