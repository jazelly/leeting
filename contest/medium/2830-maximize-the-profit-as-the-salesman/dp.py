from collections import defaultdict
from typing import List

def maximizeTheProfit(n: int, offers: List[List[int]]) -> int:
    # Initialize an array to store the maximum gold earned up to each house
    dp = [0] * (n + 1)

    d = defaultdict(list)
    for l, r, gold in offers:
        d[r+1].append((l, gold))

    for i in range(1, n + 1):
        # Initialize the maximum gold for the current house as the maximum gold from the previous house
        dp[i] = dp[i - 1]

        # Check each offer to find the maximum gold that can be earned by selling houses within the offer range
        for l, gold in d[i]:
              dp[i] = max(dp[i], dp[l] + gold)
            

    return dp[n]

print(maximizeTheProfit(5, [[1,3,2],[0,2,2],[0,0,1]]))