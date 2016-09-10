class Solution {
public:
    char findTheDifference(string s, string t) {
        char maps[256];
        for(int i =0;i<256;i++)
            maps[i] = 0;
        for(int i =0;i<s.length();i++)
            maps[s[i]]++;
        for(int i =0;i<t.length();i++){
            maps[t[i]]--;
            if (maps[t[i]]<0){
                return t[i];
            }
        }
        return 0;
    }
};