#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <vector>
#include <iostream>
#include <string>
#include <map>
#include <queue>
#include <stack>
#include <algorithm>
#include <time.h>
using namespace std;

struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {}
	
};
/** Encodes a tree to a single string. */
class Codec {
public:

	// Encodes a tree to a single string.
	string serialize(TreeNode* root) {
		if (root == NULL)
			return "nil";
		queue<TreeNode *> st;
		st.push(root);
		string ret = "";
		while (st.size()){
			auto t = st.front();
			st.pop();
			if (t == NULL)
				ret += "n|";
			else{

				ret += to_string(t->val);
				ret += "|";
				st.push(t->left);
				st.push(t->right);
			}
		}
		return ret;
	}

	// Decodes your encoded data to tree.
	TreeNode* deserialize(string data) {
		if (data == "nil")
			return NULL;
		queue<string> st;
		int before = 0;
		for (int i = 1; data[i]; i++){
			if (data[i] == '|'){
				auto tmp = data.substr(before, i-before);
				before = i+1;
				st.push(tmp);
			}
		}
		int t = atoi(st.front().c_str());
		st.pop();
		TreeNode * root = new TreeNode(t);
		queue<TreeNode *> trees[2];

		trees[0].push(root);
		int idx = 0;
		while (st.size()){
			auto tmp = trees[idx].front();
			trees[idx].pop();
			auto leftv = st.front();
			if (leftv == "n"){
				tmp->left = NULL;
			}
			else{
				TreeNode *l = new TreeNode(atoi(leftv.c_str()));
				trees[1 - idx].push(l);
				tmp->left = l;
			}
			st.pop();
			leftv = st.front();
			st.pop();
			if (leftv == "n")
				tmp->right = NULL;
			else{
				TreeNode *l = new TreeNode(atoi(leftv.c_str()));
				trees[1 - idx].push(l);
				tmp->right = l;
			}
			if (trees[idx].size() == 0)
				idx = 1 - idx;
		}
		return root;

	}
};

int main()
{
	Codec c;
	TreeNode a(1);
	TreeNode b(2);
	TreeNode dc(3);
	TreeNode d(4);
	dc.right = &d;
	a.left = &b;
	a.right = &dc;
	auto t = c.serialize(&a) ;
	auto ret = c.deserialize(t);
	cout << ret->val << endl; 
	return 0;

}