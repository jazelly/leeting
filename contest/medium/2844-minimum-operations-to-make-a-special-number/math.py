'''
https://leetcode.com/contest/weekly-contest-361/problems/minimum-operations-to-make-a-special-number/
'''
class Solution:
    def minimumOperations(self, num: str) -> int:
        n = len(num)
        nums = [int(i) for i in num]
        mem = [n] * n
        
        i = n-1
        tail = 0
        while i >= 0:
            v = nums[i]
            if v == 0:
                # lookahead for 5 or 0
                j = i-1
                dc = 0
                while j >= 0:
                    if nums[j] == 5 or nums[j] == 0:
                        break
                    j -= 1
                    dc += 1
                    
                mem[i] = dc + tail
                
            elif v == 5:
                # lookahead for 7 or 2
                j = i-1
                dc = 0
                while j >= 0:
                    if nums[j] == 7 or nums[j] == 2:
                        break
                    j -= 1
                    dc += 1
                if j >= 0:
                    mem[i] = dc + tail


            tail += 1
            i -= 1
        print(mem)
        return min(mem)
