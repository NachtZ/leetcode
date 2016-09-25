class Solution {
public:
	string toHex(int num) {
		string a[] = { "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f" };
		string res;
		if (num == 0){
			return "0";
		}
		long long int t;
		if (num <0){
			t = unsigned(num);
		}
		else{
			t = num;
		}
		while (t>0){
			res = a[t % 16] + res;
			t /= 16;
		}
		return res;
	}
};