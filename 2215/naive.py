# this beats 50% of time and 10% of space

class Solution:
    def findDifference(self, nums1, nums2):
        # distinct nums1 and nums2
        nums1 = set(nums1)
        nums2 = set(nums2)
        
        nums12 = nums1.symmetric_difference(nums2)

        nums1 = nums1.intersection(nums12)
        nums2 = nums2.intersection(nums12)
        return [nums1, nums2]