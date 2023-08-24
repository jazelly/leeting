import math
from typing import List

class Solution:
    # Extra spaces between words should be distributed as evenly as possible.
    # If the number of spaces on a line does not divide evenly between words,
    # the empty slots on the left will be assigned more spaces than the slots on the right.
    def fullJustify(self, words: List[str], maxWidth: int) -> List[str]:
        res = []
        cw = 0
        cs = 0
        l = 0
        r = 0
        while r < len(words):
            word = words[r]
            ncw = cw + len(word)
            ncs = cs
            if cw > 0:
                ncs += 1

            if ncw + ncs > maxWidth:
                line = ""

                mw = maxWidth
                while r-l-1 > 0:
                    line += words[l]

                    averageSpaceCount = (mw - cw) / (r-l-1)
                    for _ in range(math.ceil(averageSpaceCount)):
                        line += " "
                    mw -= math.ceil(averageSpaceCount)
                    mw -= len(words[l])
                    cw -= len(words[l])

                    l += 1

                line += words[l]
                l += 1
                remainSpace = maxWidth - len(line)
                for _ in range(remainSpace):
                    line += " "
                
                res.append(line)
                cw = 0
                cs = 0
                l = r
            else:
                cw = ncw
                cs = ncs
                r += 1
        
        # For the last line of text, it should be left-justified, and no extra space is inserted between words.
        lastLine = ""
        for w in words[l:]:
            lastLine += w
            lastLine += " "

        lastLine = lastLine[:-1]
        remainSpace = maxWidth - len(lastLine)
        for _ in range(remainSpace):
            lastLine += " "

        res.append(lastLine)
        return res




s = Solution()
words = ["This", "is", "an", "example", "of", "text", "justification."]
maxWidth = 16
s.fullJustify(words, maxWidth)