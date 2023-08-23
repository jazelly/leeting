from collections import Counter
import heapq

class Solution:
    def reorganizeString(self, s: str) -> str:
        res = ""
        def max_heap_with_occurrences(s):
            char_count = Counter(s)
            max_heap = [(-count, char) for char, count in char_count.items()]
            heapq.heapify(max_heap)
            return max_heap

        heap = max_heap_with_occurrences(s)

        while heap:
            count, c1 = heapq.heappop(heap)
            new_count = -count - 1

            res += c1

            if not heap:
                if -count > 1:
                    return ""
                return res
            count2, c2 = heapq.heappop(heap)
            res += c2
            newCount2 = -count2-1
            if newCount2 > 0:
                heapq.heappush(heap, (-newCount2, c2))

            if new_count > 0:
                heapq.heappush(heap, (-new_count, c1))


        return res