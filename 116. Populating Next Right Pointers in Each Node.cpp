class Solution {
public
void connect(TreeLinkNode root) {
		if (root == NULL  root -left == NULL) return;
		root-left-next = root-right;
		if (root-next != NULL)
			root-right-next = root-next-left;
		if (root-left != NULL && root-left-left != NULL)
		{
			connect(root-left);
			connect(root-right);
		}
    
}
};