struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB) {
    		int lengthA = 0, lengthB = 0;
		struct ListNode * temp = headA;
		while (temp != NULL)
		{
			temp = temp->next;
			lengthA++;
		}
		temp = headB;
		while (temp != NULL)
		{
			temp = temp->next;
			lengthB++;
		}
		int n = 0;
		if (lengthA > lengthB)
		{
			n = lengthA - lengthB;
			while (n--)
				headA = headA->next;
		}
		else
		{
			n = lengthB - lengthA;
			while (n--)
				headB = headB->next;
		}
		while (1)
		{
			if (headA == headB)
				return headA;
			headA = headA->next;
			headB = headB->next;
		}
	}
