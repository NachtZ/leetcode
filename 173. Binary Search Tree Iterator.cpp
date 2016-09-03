/**
 * Definition for binary tree
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class BSTIterator {
public:
    stack<TreeNode *> s;
    BSTIterator(TreeNode *root) {
        pushLeft(root);
    }

    /** @return whether we have a next smallest number */
    bool hasNext() {
        return !s.empty();
    }

    /** @return the next smallest number */
    int next() {
        TreeNode* top = s.top();
        s.pop();
        pushLeft(top->right);
        return top->val;
    }
    
    void pushLeft(TreeNode* root)
    {
        if(root != NULL)
        {
            s.push(root);
            TreeNode* cur = root;
            while(cur->left)
            {
                s.push(cur->left);
                cur = cur->left;
            }
        }
    }
};
/**
 * Your BSTIterator will be called like this:
 * BSTIterator i = BSTIterator(root);
 * while (i.hasNext()) cout << i.next();
 */