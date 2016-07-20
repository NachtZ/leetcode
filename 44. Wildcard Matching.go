
func isMatch(s string, p string) bool {
       ls,lp := len(s),len(p)
       res := [][]bool{}
       for i:=0;i<ls+1;i++{
           rt := make([]bool,lp+1)
           res = append(res,rt)
       }
       res[0][0] = true
       for i:=1;i<=lp;i++{
           if p[i-1] == '*'{
               res[0][i] = res[0][i-1]
           }
       }
       for i:=1;i<=ls;i++{
           for j:=1;j<=lp;j++{
               if p[j-1]!='*'{
                   res[i][j] = (s[i-1] == p[j-1] || p[j-1] == '?') && res[i-1][j-1]
         
               }else{
                   res[i][j] = res[i-1][j]||res[i][j-1]
               }
           }
       }
       return res[ls][lp]
}