func calcEquation(equations [][]string, values []float64, query [][]string) []float64 {
    m := make(map[[2]string]float64)
	//tm := make(map[string]float64)
    for i:=0;i<len(equations);i++{
		a,b :=equations[i][0],equations[i][1]
		m[[2]string{a,b}] =values[i]
		m[[2]string{b,a}] = 1/values[i]
		m[[2]string{a,a}] =1.0
		m[[2]string{b,b}] =1.0
		for j :=i+1;j<len(equations);j++{
			//in here to construct the map
			c,d := equations[j][0],equations[j][1]
			if a!=c &&a!=d && b!=c && b!=d{
				continue// can't get any new info
			}
			if a == c &&b == d || a == d&& b == c{
				continue// can't get any new info
			}
			if a == c{
				m[[2]string{d,b}] = values[i]/values[j]
				m[[2]string{b,d}] = values[j]/values[i]
			}
			if a == d{
				m[[2]string{c,b}] = values[i]*values[j]
				m[[2]string{b,c}] = 1/(values[i]*values[j])
			}
			if b == c{
				m[[2]string{a,d}] = values[i]*values[j]
				m[[2]string{d,a}] = 1/(values[j]*values[i])				
			}
			if b == d{
				m[[2]string{a,c}] = values[i]/values[j]
				m[[2]string{c,a}] = values[j]/values[i]						
			}
		}
	}
    for	k,v := range m{
		a,b := k[0],k[1]
		m[[2]string{a,b}] =v
		m[[2]string{b,a}] = 1/v
		m[[2]string{a,a}] =1.0
		m[[2]string{b,b}] =1.0
		for k1,v1 := range m{
			//in here to construct the map
			c,d := k1[0],k1[1]
			if a!=c &&a!=d && b!=c && b!=d{
				continue// can't get any new info
			}
			if a == c &&b == d || a == d&& b == c{
				continue// can't get any new info
			}
			if a == c{
				m[[2]string{d,b}] = v/v1
				m[[2]string{b,d}] = v1/v
			}
			if a == d{
				m[[2]string{c,b}] = v*v1
				m[[2]string{b,c}] = 1/(v*v1)
			}
			if b == c{
				m[[2]string{a,d}] = v*v1
				m[[2]string{d,a}] = 1/(v*v1)				
			}
			if b == d{
				m[[2]string{a,c}] = v/v1
				m[[2]string{c,a}] = v1/v					
			}
		}
	}
	tmp := [2]string{"",""}
	var ret []float64
	for i:=0;i<len(query);i++{
		tmp[0],tmp[1] = query[i][0],query[i][1]
		v, ok := m[tmp]
		if ok{
			ret = append(ret,v)
		}else{
			ret = append(ret,-1.0)
		}
	}
	return ret

}