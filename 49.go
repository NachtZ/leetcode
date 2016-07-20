

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}


func groupAnagrams(strs []string) [][]string {
    ret := [][]string{}
    dic := make(map[string]int,len(strs))
    idx := 0
    for _,str := range strs{
        tmp :=SortString(str)        
        val,ok := dic[tmp]
        if ok == false{
            dic[tmp] = idx
        
            ret = append(ret,[]string{})
            ret[idx] = append(ret[idx],str)
            idx ++
        }else{
            ret[val] = append(ret[val],str)
        }
    } 
    return ret  
}
```