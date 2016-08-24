/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 * 当环存在的时候，链表有一个性质，就是头结点到入环点和快慢指针碰撞点到入环点的距离相等。利用这个性质，我们就能很方便的求出入环点。只要从碰  *撞点和头结点一起走，知道两个相遇就是入环点。
 */
class Solution {
public:
	ListNode *detectCycle(ListNode *head) {
		if (head == NULL || NULL == head->next)
			return NULL;
		ListNode *fast, *slow;
		fast = head -> next->next;
		slow = head->next;
		while (fast != NULL && fast->next != NULL)
		{
			if (fast == slow)
				break;
			fast = fast->next->next;
			slow = slow->next;
		}
		if (fast == NULL || fast->next == NULL)
			return NULL;
		while (fast != head)
		{
			fast = fast->next;
			head = head->next;
		}
		return head;
	}
};