/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int getLeft(TreeNode * root){
        int count = 0;
        while(root->left !=NULL){
            ++count;
            root = root ->left;
        }
        return count;
    }
    int getRight(TreeNode * root){
        int count = 0;
        while(root->right !=NULL){
            ++count;
            root = root ->right;
        }
        return count;
    }
    int countNodes(TreeNode* root) {
        if(root == NULL)return 0;
        int l = getLeft(root)+1;
        int r = getRight(root) +1;
        if(l == r)return (2<<(l-1))-1;
        else
        return countNodes(root->left)+countNodes(root->right)+1;
        
    }
};