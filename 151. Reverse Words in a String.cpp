class Solution {
public:
void reverseWords(string &s) {
       string temp="",result="";
       for(int i=0;i<s.length();i++)
       if(*(s.c_str()+i)==' '){ 
           result= temp + (temp==""||result=="" ? "": " ")+result;
           temp="";
       }
       else
        temp=temp+*(s.c_str()+i); 
       s= temp + (temp==""||result=="" ? "": " ")+result;
}
};