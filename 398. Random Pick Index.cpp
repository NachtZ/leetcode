class Solution {
	vector<int> num;
public:
	Solution(vector<int> nums) {
		num = nums;
	}

	int pick(int target) {
		int index = -1;
		int count = 0;
		for (int i = 0; i < num.size(); i++){
			if (num[i] == target){
				count++;
				int r = rand() % count + 1;
				if (r == count){
					index = i;
				}
			}
		}
		return index;
	}
};