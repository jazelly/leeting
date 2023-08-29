class Solution:
    def bestClosingTime(self, customers: str) -> int:
        left = defaultdict(int)
        right = defaultdict(int)

        for c in customers:
            right[c] += 1

        res = [right['Y']]
        i = 0
        while i < len(customers):
            right[customers[i]] -= 1
            left[customers[i]] +=1
            res.append(left['N'] + right['Y'])
            i += 1
            

        target = min(res)
        i = 0
        for r in res:
            if r == target:
                break
            i += 1

        return i