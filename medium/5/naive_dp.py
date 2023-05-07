# very slow
# only beats 5% time and space

class Solution:
    def longestPalindrome(self, s: str) -> str:       
        if len(s) == 0:
            return ''
        if len(s) == 1:
            return s[0]

        dp = []
        
        for i in range(1, len(s)):
            if i-2 >= 0 and s[i-2] == s[i]:
                dp.append((i-2, i))

            elif s[i] == s[i-1]:
                dp.append((i-1, i))
        
        if len(dp) > 0:
            re = dp[0][1]
            rs = dp[0][0] 
            l = len(dp)
            i = 0
            while i < l:
                item = dp[i]
                st = item[0]
                e = item[1]
                if e - st > re - rs:
                    rs = st
                    re = e

                if st - 1 >= 0 and e + 1 < len(s) and s[st -1] == s[e + 1]:
                    dp.append((st - 1, e + 1))
                    l += 1
                elif st - 1 >= 0 and s[st -1] == s[e] and len(set(s[st-1:e+1])) == 1:
                    dp.append((st - 1, e))
                    l += 1
                i += 1

            return s[rs:re+1]
        else:
            return s[0]