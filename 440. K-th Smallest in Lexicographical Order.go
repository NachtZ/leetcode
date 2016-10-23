var Len,n,m,ret int
var num,ans [20]int

func valueOfBegin(b int ) int{
	cnt :=0
	for i:=b;i<=Len;i++{
		cnt = cnt *10 + num[Len-i+1]
	}
	return cnt
}
func countofNum(a [20]int,anslen int) int{
	k,cnt :=1,0
	for i:=0;i+anslen <Len;i++{
		cnt +=k
		k*=10
	}
	op :=0
	for i:=1;i<=anslen;i++{
		id := Len-i+1
		if num[id]>a[i]{
			op = -1
			break
		}else if num[id] < a[i]{
			op = 1
			break
		}
	}
	if op == 0{
		if anslen == Len{
			return 1
		}
		return valueOfBegin(anslen+1)+1+cnt
	}
	if op == 1{
		return cnt
	}
		k = 1
		for i:=0;i<Len-anslen;i++{
			k*=10

		}
		return cnt +k
}

func dfs(ans [20]int, cur,m int){
	if m == 1{
		ret = 0
		for i:=1;i<=cur;i++{
			ret = ret *10 + ans[i]
		}
		return
	}
	m --
	for i:=0;i<=9;i++{
		ans[cur+1] = i
		cnt := countofNum(ans,cur+1)
		if cnt >=m{
			dfs(ans,cur+1,m)
			return
		}else{
			m -= cnt
		}
	}
}

func findKthNumber(a int, k int) int {
    Len = 0
	n = a
	m = k
	for n>0{
		Len ++ 
		num[Len] = n%10
		n/=10
	}
	for i:=1;i<=9;i++{
		ans[1] = i
		cnt := countofNum(ans,1)
		if cnt >=m{
			dfs(ans,1,m)
			break
		}else{
			m -= cnt
		}
	}
	return ret
}