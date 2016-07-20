
func strStr(h string, n string) int {
    for i:=0;;i++{
        for j:=0;;j++{
            if j == len(n){
                return i
            }
            if i+j == len(h){
                return -1
            }
            if (h[i+j]!=n[j]){
                break
                
            }
        }
    }
}