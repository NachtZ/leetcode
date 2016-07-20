# Definition for singly-linked list.
 class ListNode(object):
     def __init__(self, x):
         self.val = x
         self.next = None

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        num1 = 0
        num2 = 0
        ten = 1
        while l1 != None:
            num1 += l1.val*ten
            ten *= 10
            l1 = l1.next
        ten = 1
        while l2 != None:
            num2 += l2.val*ten
            ten *=10
            l2 = l2.next
        num3 = num1 +num2
        t = num3%10
        l3 = ListNode(t)
        head = l3
        num3 = num3/10
        while num3 != 0:
            t = num3%10
            l3.next = ListNode(t)
            l3 = l3.next
            num3/=10
        return head

s = Solution()
p
            
