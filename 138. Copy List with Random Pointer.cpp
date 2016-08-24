class Solution {
public:

    RandomListNode *copyRandomList(RandomListNode *head) {
        //unordered_map < RandomListNode *, RandomListNode * > m;
        if (head == NULL)return NULL;
        unordered_map < int , RandomListNode * > m;
        RandomListNode * tmp = head -> next,*tmp1;
        RandomListNode *dst = new RandomListNode(head->label);
        m[head->label] = dst;
        dst->next = NULL;
        dst->random = NULL;
        tmp1 = dst;
        while (tmp){
            RandomListNode * node = new RandomListNode(tmp->label);
            node->random = tmp->random;
            tmp1->next = node;
            tmp1 = tmp1->next;
            m[tmp->label] = node;
            tmp = tmp->next;
        }
        tmp1->next = NULL;
        tmp1 = dst;
        tmp = head;
        while (tmp1){
            if (tmp->random)
                tmp1->random = m[tmp->random->label];
            else
                tmp1->random = NULL;
            tmp1 = tmp1->next;
            tmp = tmp->next;
        }
        return dst;
    }
};