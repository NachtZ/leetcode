class Solution {
public:
    int minimumTotal(vector<vector<int>>& t) {
        if(t.size() == 0)return 0;
        if(t.size()==1)return t[0][0];
        int ** dp = new int*[t.size()];
        dp[0] = new int[t.size()*t.size()+1];
        for(int i =1;i!=t.size();++i){
            dp[i] = dp[i-1]+ t.size();
        }
        dp[0][0] = t[0][0];
        for(int i =1;i!=t.size();++i){
            dp[i][0] = dp[i-1][0]+t[i][0];
            dp[i][i] = dp[i-1][i-1]+t[i][i];
            for(int j =1;j!=i;++j){
                dp[i][j] = t[i][j]+(dp[i-1][j-1]>dp[i-1][j]?dp[i-1][j]:dp[i-1][j-1]);
            }
        }
        int min = dp[t.size()-1][0];
        for(int i =1;i!=t.size();++i){
            if(min>dp[t.size()-1][i])min = dp[t.size()-1][i];
        }
        return min;
    }
};