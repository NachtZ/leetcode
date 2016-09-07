class Solution {
public:
    vector<string>ans;
    void dfs(int cur, long long cnum, long long num, long long res, string str, string s){
        cnum = cnum * 10 + s[cur] - '0';
        if (cur+1 == s.size()){
            if (res + num*cnum == 0)ans.push_back(str + s[cur]);
            return;
        }
        if (cnum != 0)dfs(cur + 1, cnum, num, res, str + s[cur], s);//no op here and no leading 0
        dfs(cur + 1, 0, 1, res + num*cnum, str + s[cur] + '+', s);//+
        dfs(cur + 1, 0, -1, res + num*cnum, str + s[cur] + '-', s);//-
        dfs(cur + 1, 0, num*cnum, res, str + s[cur] + '*', s);//* 
    }
    vector<string> addOperators(string num, int target) {
        if (num.size() > 0)dfs(0, 0, 1, -(long long)target, "", num);
        return ans;
    }
};