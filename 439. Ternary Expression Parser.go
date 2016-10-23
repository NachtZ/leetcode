func parseTernary(e string) string {
    if e[0] != 'T' && e[0]!='F'{
		return e
	}
	c1,c2,i := 0,0,1
	for i=1;i<len(e);i++{
		if e[i] == '?'{
			c1 ++
		}else{
			if e[i] == ':'{
				c2 ++
				if c1 == c2{
					break;
				}
			}
		}
	}
	if i == len(e){
		return e
	}
	if e[0] == 'T'{
		return parseTernary(e[2:i])
	}
	return parseTernary(e[i+1:])
}