class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if(root == NULL || p == NULL || q == NULL)return NULL;
        //swap p and q.
        if(q->val == root->val || p->val == root->val)return root;
        if(p->val > q->val){
            TreeNode * tmp = p;
            p = q;
            q = tmp;
        }
        if(p->val < root->val && q->val > root->val)
            return root;
        if(p->val > root->val)
            return lowestCommonAncestor(root->right,p,q);
        if(q->val < root->val)
            return lowestCommonAncestor(root->left,p,q);
        return NULL;
    }
};