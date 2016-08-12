class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        t = x
        ret = 0
        if x <0:
            x *=-1
        while x !=0:
            ret = ret *10 + x%10
            x = int(x/10)
  
        if ret >= 2147483648:
            return 0
        if t <0:
            ret *= -1 
        return ret