class Solution {
public:
	TreeLinkNode *findNext(TreeLinkNode *root)
	{
		if (!root)return NULL;
		if (root->left) return root->left;
		if (root->right) return root->right;
		return findNext(root->next);
	}

	void connect(TreeLinkNode *root) {
		if (!root) return ;
		TreeLinkNode *next = findNext(root);
		while (root)
		{
			if (root->left) root->left ->next = root->right ? root->right : findNext(root ->next);
			if (root->right) root->right ->next = findNext(root->next);
			root = root->next;
			if(!root)
			{
				root = next;
				next = findNext(root);
			}
		}
	}
};