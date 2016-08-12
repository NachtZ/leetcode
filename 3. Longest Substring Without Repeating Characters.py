
class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        ch = {}
        for i in s:
            ch[i] = -1
        index = -1
        max=0
        for i in range(len(s)):
            if ch[s[i]] > index:
                index = ch[s[i]]
            if max < i - index:
                max = i -index
            ch[s[i]] = i
        
        return max

s = Solution()
print(s.lengthOfLongestSubstring("abcabcbb"))