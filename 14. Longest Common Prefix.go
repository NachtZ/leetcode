

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0{
        return ""
    }
    if len(strs) == 1{
        return strs[0]
    }
    pos := len(strs[0]) -1
    for i :=1;i<len(strs);i++{
        var j int
        if pos > len(strs[i]) -1{
            pos = len(strs[i]) -1
        }
        for j=0;j<=pos && j<=len(strs[i])-1;j++{
            if strs[0][j]!=strs[i][j]{
                pos = j-1
                break
            }
        }
    }
    if(pos == -1){
        return ""
    }else {
        return strs[0][0:pos+1]
    }
}