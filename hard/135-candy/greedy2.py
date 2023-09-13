from typing import List


class Solution:
    def candy(self, ratings: List[int]) -> int:
        n = len(ratings)
        
        candies = [1] * n
        if n == 1:
            return sum(candies)
            
        index = 0
        while index <= n-1:
            lc = candies[index]
            rc = candies[index]
            if index-1 >= 0 and ratings[index-1] < ratings[index]:
                lc = candies[index-1] + 1
            if index+1 <= n-1 and ratings[index] > ratings[index+1]:
                rc = candies[index+1] + 1
            candies[index] = max(lc, rc)
            index += 1

        index = n-2
        while index >= 0:
            lc = candies[index]
            rc = candies[index]
            if index+1 <= n-1 and ratings[index] > ratings[index+1]:
                rc = candies[index+1] + 1
            candies[index] = max(lc, rc)
            index -= 1


        return sum(candies)