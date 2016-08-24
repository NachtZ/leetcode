class Solution {
public:
    unordered_map<string, vector<string>> dp;
    vector<string> d(const string cur, const unordered_set<string>& wordDict){
        if(dp.find(cur) != dp.end()){
            return dp[cur];
        }
        vector<string> ans;
        if(wordDict.find(cur) != wordDict.end()){
            ans.push_back(cur);
        }
        // each time , remove a word from end
        for(int i = cur.length() - 1;i > 0; --i){
            string last = cur.substr(i,cur.length() - i);
            if(wordDict.find(last) != wordDict.end()){
                vector<string> pre = d(cur.substr(0, i), wordDict);
                for(int i = 0;i < pre.size(); ++i){
                    pre[i] += " " + last;
                }
                ans.insert(ans.end(), pre.begin(), pre.end());
            }
        }
        dp[cur] = ans;
        return ans;
    }
    vector<string> wordBreak(string s, unordered_set<string>& wordDict) {
        return d(s, wordDict);
    }
};