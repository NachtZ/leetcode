class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        ptr1 = 0
        ptr2 = 0
        size = len(nums1) + len(nums2)
        if size == 0 :
            return 0
        
        middle = int(size/2)
        nums = []
        for i in range(middle+1):
            if ptr1 == len(nums1) and ptr2 != len(nums2):
                nums.append(nums2[ptr2])
                ptr2 +=1
                continue
            if ptr1 != len(nums1) and ptr2 == len(nums2):
                nums.append(nums1[ptr1])
                ptr1 +=1
                continue
            if nums1[ptr1] <= nums2[ptr2]:
                nums.append(nums1[ptr1])
                ptr1 +=1
            else:
                nums.append(nums2[ptr2])
                ptr2 +=1
        
        if size % 2 == 0:
            return (nums[middle-1]+nums[middle])/2.0
        else:
            return nums[middle]

s = Solution()
print(s.findMedianSortedArrays([1,3],[2]))