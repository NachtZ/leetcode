class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if len(s) <=1 or numRows <=1:
            return s
    
        str = ""
        size = 2*numRows -2
        for i in range(numRows):
            for j in range(i,len(s),size):
                str += s[j]
                if i !=0 and i!= numRows -1:
                    tmp = j + size -2*i
                    if tmp < len(s):
                        str += s[tmp]
        return str

s = Solution()
print(s.convert("PAYPALISHIRING", 3))