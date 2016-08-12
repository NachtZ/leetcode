class Solution {
 public:
	 int myAtoi(string str) {
		 long long int res = 0;
		 int pos = 1;
		 int i = 0;
		 for (i = 0; str[i] && str[i] == ' ';i++);
		 if (str[i] == '+')
		 {
			 pos = 1;
			 i++;
		 }
		 else if (str[i] == '-')
		 {
			 i ++;
			 pos = -1;
		 }
		 for (; i < str.size(); i++)
		 {
				 if (pos *res >= INT_MAX)return INT_MAX;
		 if (pos *res <= INT_MIN)return INT_MIN;
			 if (str[i]>'9' || str[i] < '0')return res*pos;
			 res = res * 10 + str[i] - '0';
		 }
				 if (pos *res >= INT_MAX)return INT_MAX;
		 if (pos *res <= INT_MIN)return INT_MIN;
		 if (pos == -1)res *= -1;
		 return res;
	 }
 };