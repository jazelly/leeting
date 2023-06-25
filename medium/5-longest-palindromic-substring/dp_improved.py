# this beats 41% time and 20% space

class Solution:
    def longestPalindrome(self, s: str) -> str:       
        if len(s) == 0:
            return ''
        if len(s) == 1:
            return s[0]

        dp_odd = []
        dp_even = []
        
        # find all the starting point
        # O(n)
        for i in range(1, len(s)):
            if i-2 >= 0 and s[i-2] == s[i]:
                dp_odd.append(i-1)

            if s[i] == s[i-1]:
                dp_even.append((i-1, i))

        # short circuit if there is no starting point
        if len(dp_odd) == 0 and len(dp_even) == 0:
            return s[0]
        
        re = -1
        rs = -1
        print(dp_odd)
        if len(dp_odd) > 0:
            rs = dp_odd[0] - 1
            re = dp_odd[0] + 1
            for i in range(len(dp_odd)):
                st = dp_odd[i] - 1
                e = dp_odd[i] + 1
                # expanding
                while st - 1 >= 0 and e + 1 < len(s) and s[st - 1] == s[e + 1]:
                    st -= 1
                    e += 1
                    if e - st > re - rs:
                        rs = st
                        re = e

        print('--------')
        print(dp_even)
        if len(dp_even) > 0:
            # if no solutions found in dp_odd
            if re < 0:
                re = dp_even[0][1]
                rs = dp_even[0][0]

            for i in range(len(dp_even)):
                st = dp_even[i][0]
                e = dp_even[i][1]
                while st - 1 >= 0 and e + 1 < len(s) and s[e + 1] == s[st - 1]:
                    st -= 1
                    e += 1
                    dp_even[i] = (st, e)
                    if e - st > re - rs:
                        rs = st
                        re = e
        print(dp_even)

        return s[rs:re+1]