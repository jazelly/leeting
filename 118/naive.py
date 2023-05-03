class Solution:
    def generate(self, numRows: int):
        # based on a given list, get next list
        def f(l):
            nl = [l[i] + l[i+1] for i in range(len(l) - 1)]
            return [1] + nl + [1]
        
        res = [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]
        currentList = res[4]
        if numRows > 5:
            for i in range(5, numRows):
                currentList = f(currentList)
                res.append(currentList)
        return res[0:numRows]