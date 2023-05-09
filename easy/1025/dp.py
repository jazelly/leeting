# https://leetcode.com/problems/divisor-game/description/

class Solution:
    def divisorGame(self, n: int) -> bool: 
        # dp stores the optimal remaining steps to finish game
        # however, it is not the lowest steps
        # Alice is seeking for an even number of steps, so that Bob cannot win
        dp = [2 * n for i in range(n+1)]

        dp[0] = 0
        dp[1] = 0

        # get all the solution for i in n
        for i in range(2, n+1):
            # find the solutions
            solutions = [dp[i-j] for j in range(1, i) if i % j == 0]
            if len(solutions) == 0:
                dp[i] = 0
                continue
            print(solutions)

            for s in solutions:
                if s % 2 == 0:
                    dp[i] = s + 1
        print(dp)
        # as long as the optimal steps to finish the game is not even, Alice wins
        return dp[-1] % 2 != 0
