type P [][]int 
func(p P) Len()int{
	return len(p)
}
func(p P)Swap(i,j int){
	p[i][0],p[i][1],p[j][0],p[j][1] = p[j][0],p[j][1],p[i][0],p[i][1]
}
func (p P)Less(i,j int)bool{
	if p[i][1]!=p[j][1]{
		return p[i][1] < p[j][1]
	}
	return p[i][0] < p[j][0]
}
//1.sort
//2.取当前队列
func reconstructQueue(people [][]int) [][]int {
	res := make([][]int,len(people))
	for i:=0;i<len(res);i++{
		res[i]  = make([]int,2)
	}
	n := len(people)
	for  i:=0;i<n;i++{
		
		sort.Sort(P(people))
		res[i][0],res[i][1] = people[0][0],people[0][1]
		for j:=1;j<len(people);j++{
			if people[0][0] >=people[j][0]{
				people[j][1] --;
			}
		}
		people = people[1:]
	}
	for i :=1;i<len(res);i++ {
		count :=0
		for j :=0;j<i;j++ {
			if res[i][0] <= res[j][0]{
				count ++;
			}
		}
		res[i][1] = count
	}
	return res;
}