class Solution {
public:
vector<vector<string> > partition(string s) 
    {
        vector<vector<string> > rs;
        vector<string> tmp;
        storePatition(rs, tmp, s);
        return rs;
    }
 
    void storePatition(vector<vector<string> > &rs, vector<string> &tmp,
        string &s)
    {
        if (s.empty())
        {
            rs.push_back(tmp);
            return;
        }
        for (int i = 1; i <= s.length(); i++)
        {
            string a = s.substr(0, i);
            if (isPalindrome(a))
            {
                tmp.push_back(a);
                string b = s.substr(i);
                storePatition(rs, tmp, b);
                tmp.pop_back();
            }
        }
    }
 
    bool isPalindrome(string &s)
    {
        for (int i = 0, j = s.length()-1; i < j; i++, j--)
        {
            if (s[i] != s[j]) return false;
        }
        return true;
    }
};