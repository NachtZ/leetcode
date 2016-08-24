class Solution {
public:
    bool hasCycle(ListNode *head) {
        if(head == NULL||NULL == head->next)
            return false;
        ListNode *fast,*slow;
        fast = head->next->next;
        slow = head->next;
        while(fast!=NULL && fast->next != NULL)
        {
            if(fast == slow)
                return true;
            fast = fast ->next ->next;
            slow = slow -> next;
        }
        return false;
    }
};