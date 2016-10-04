#include<stdio.h>
#include<unordered_map>
#include<vector>
#include<stdlib.h>
#include<time.h>
#include<set>

#include<stdlib.h>
#include<string>
#include<iostream>
using namespace std;
class Solution {
public:
	int longestPalindrome(string s) {
		int maps[128];
		bool hasS = false;
		int count = 0;
		for (int i = 0; i < 128; i++){
			maps[i] = 0;
		}
		for (auto c : s){
			maps[c] ++;
		}
		for (int i = 0; i < 128; i++){
			if (maps[i] & 1){
				hasS = true;
			}
			count += (maps[i] >> 1);
		}
		count <<= 1;
		if (hasS)
			count++;
		return count;
	}
};
int main()
{
	Solution s;
	cout << s.longestPalindrome("abccccdd") << endl;
	return 0;
}
