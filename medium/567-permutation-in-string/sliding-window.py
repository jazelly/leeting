class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        w = defaultdict(int)
        need = defaultdict(int)
        for c in s1:
            need[c] += 1

        l = 0
        r = 0
        valid = 0
        while r < len(s2):
            if s2[r] in need:
                w[s2[r]] += 1
                if w[s2[r]] == need[s2[r]]:
                    valid += 1
                elif w[s2[r]] > need[s2[r]]:
                    # move left to a point that w[r] is valid
                    while l < r:
                        tl = s2[l]
                        w[tl] -= 1
                        l+=1
                        if tl == s2[r]:
                            break
                    
                    # recalculate valid
                    valid = 0
                    for k in w:
                        if w[k] == need[k] and k in need:
                            valid += 1
                r+=1
            else:
                w = defaultdict(int)
                valid = 0
                r += 1
                l = r
            if valid == len(need):
                return True
        return False
