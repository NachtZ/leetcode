bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int l = 1;
        int r = n;
        if(isBadVersion(1) == true)return 1;
        int mid;
        while(1){
            mid = l + (r-l)/2;
            if(mid == l)return r;
            if(isBadVersion(mid)){
                r = mid;
            }
            else{
                l = mid;
            }
        }

    }
};