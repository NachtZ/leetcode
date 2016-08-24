class Solution {
public:
    bool wordBreak(string s, unordered_set<string>& wordDict) {
        bool * dp = new bool[s.length()+1];
        for(int i =0;i!=s.length()+1;++i)
            dp[i] =false;
        dp[0] = true;
        for(int i =0;i!=s.length();++i){
            if(dp[i]){
                for(int j=1;j+i<=s.length();++j){
                    if(wordDict.count(s.substr(i,j))>0)
                        dp[i+j] = true;
                }
            }
        }
        return dp[s.length()];
    }
};