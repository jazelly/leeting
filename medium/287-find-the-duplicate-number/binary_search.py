class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        n = len(nums)
        low = 1
        high = n - 1
        while low < high:
            mid = low + (high - low) // 2
            cnt = 0
            for i in range(n):
                if nums[i] <= mid:
                    cnt += 1
            
            if cnt <= mid:
                low = mid + 1
            else:
                high = mid

        return low

