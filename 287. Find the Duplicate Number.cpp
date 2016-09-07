class Solution {
public:
    int findDuplicate(vector<int>& nums) {
        int left = 1;
        int right = nums.size()-1;
        int mid = left+(right - left)/2;
        while(left < right){
            int c = 0;
            mid = left+(right - left)/2;
            for(int i =0;i< nums.size();i++){
                if(nums[i] <=mid)c ++;
            }
            if(c >mid)right = mid;
            else
                left = mid+1;
        }
        return left;
    }
};