class Solution {
public:
    int min(int a, int b){
        if (a>b)
            return b;
        return a;
    }
    int numSquares(int n) {
        int *dp = new int[n+1];
        for(int i =0;i!=n+1;++i)dp[i] = 4;
        for(int i =0;i*i<=n;++i){
            dp[i*i] = 1;
        }
        for(int a = 0;a!=n+1;++a){
            for(int b = 0;b*b<=a;++b){
                dp[a] = min(dp[a],dp[a-b*b]+1);
            }
        }
        return dp[n];
    }
};