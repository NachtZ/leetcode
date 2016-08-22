#include<io.h>
#include<iostream>
#include<vector>
#include<string>
using namespace std;
class Solution {
public:
	vector<int> ret;
	int max;
	void help(int sum){
		if (sum >= max){
			return;
		}
		ret.push_back(sum);
		for (int i = 0; i <= 9; i++){
			help(sum * 10 + i);
		}
	}
	vector<int> lexicalOrder(int n) {
		max = n;
		for (int i = 1; n>=i && i <= 9; i++){
			help(i);
		}
		return ret;
	}
};

int main(){
	Solution s;
	vector<int> a = s.lexicalOrder(13);
	for (int i = 0; i < a.size(); i++){
		cout << a[i] << endl;
	}

}
