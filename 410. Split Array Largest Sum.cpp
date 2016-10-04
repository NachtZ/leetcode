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
	int splitArray(vector<int>& nums, int m) {
		int max = 0,sum = 0;
		for (auto num : nums){
			if (num > max)
				max = num;
			sum += num;
		}
		if (m == 1){
			return sum;
		}
		long l = max, r = sum;
		while (l <= r){
			long mid = l + (r - l) / 2;
			if (can(mid, nums, m))
				r = mid - 1;
			else
				l = mid + 1;
		}
		return l;
	}
	bool can(long target, vector<int>& nums, int m){
		int count = 1;
		long total = 0;
		for (auto num : nums){
			total += num;
			if (total > target){
				total = num;
				count++;
				if (count > m){
					return false;
				}
			}
		}
		return true;
	}
};
int main()
{
	Solution s;
	cout << s.longestPalindrome("abccccdd") << endl;
	return 0;
}
