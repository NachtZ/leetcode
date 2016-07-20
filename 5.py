class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        l = len(s)
        str = "#"
        for i in s:
            str += i + "#"
        str = "^" + str + "$"
        p = [0]*len(str)
        C = 0
        R = 0
        for i in range(len(str)-1):
            i_mirror = C - (i -C)
            diff = R -i
            if diff >=0:
                if p[i_mirror] <diff:
                    p[i] = p[i_mirror]
                else:
                    p[i] = diff
                    while str[i+p[i]+1] == str[i-p[i]-1]:
                        p[i] += 1
                    C = i
                    R = i + p[i]
            else:
                p[i] = 0
                while str[i+p[i]+1] == str[i-p[i]-1]:
                    p[i] += 1
                C = i
                R = i + p[i]
        maxLen = 0
        centerIndex = 0
        for i in range(len(str)-1):
            if p[i]>maxLen:
                maxLen = p[i]
                centerIndex = i
        
        return s[int((centerIndex-1-maxLen)/2):int((centerIndex-1-maxLen)/2+maxLen)]