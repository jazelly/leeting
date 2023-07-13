# beats 5% on space and time
class Solution:
    def generateParenthesis(self, n):
        def f(left, right, s):
            print('In stack: ', left, right)
            if len(s) == 2 * n:
                print('satisfied')
                res.append(s) 
            
            if left < n:
                print('pushing to stack: ', left + 1, right)
                f(left + 1, right, s + '(')
                print('popped from stack with context: ', left, right)
                print('when left < n')
            if right < left:
                print('pushing to stack: ', left, right + 1)
                f(left, right + 1, s + ')')
                print('popped from stack with context: ', left, right)
                print('when right < left')

        res = []
        f(0, 0, '')
        return res
