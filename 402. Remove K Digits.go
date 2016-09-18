var res string
func help(num string, k int) {
    if k <= 0{
		res = res + num
		return
	}
	if k >= len(num){
		return ;
	}
	minIdx := 0
	for i:=1;i<=k;i++{
		if num[i] < num[minIdx]{
			minIdx = i
		}
	}

	res = res + string(num[minIdx])
	new := num[minIdx+1:]
	help(new,k - minIdx)

}

func removeKdigits(num string, k int) string {
    res = ""
	help(num,k)
	i:=0
	for i=0;i<len(res)&&res[i]=='0';i++{
	}
	if i<len(res){
		res = res[i:]
	}else{
		res = "0"
	}
	return res
}