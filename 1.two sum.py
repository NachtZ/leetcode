import copy

class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        cp = copy.deepcopy(nums)
        cp.sort()
        p1 = 0
        p2= len(nums) -1
        while p1 < p2:
            if cp[p1] +cp[p2] == target:
                break
            elif cp[p1]+cp[p2]>target:
                p2 = p2-1
            else:
                p1 = p1 +1
        ans = []
        for i in range(len(nums)):
            if p1 != -1 and cp[p1] == nums[i]:
                p1 = -1
                ans.append(i)
                continue
            if p2 !=-1 and cp[p2] == nums[i]:
                p2 = -1
                ans.append(i)
                continue
            if len(ans) == 2:
                return ans
        
        return ans

s = Solution()
print(s.twoSum([3,2,4],6))
