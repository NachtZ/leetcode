func addStrings(num1 string, num2 string) string {
    l1,l2 := len(num1),len(num2)
	if l1 == 0{
		return num2
	}
	if l2 == 0{
		return num1
	}
	ret := ""
	add :=0
	for l1,l2 = l1-1,l2-1 ;l1>=0 && l2 >=0;l1,l2 = l1-1,l2-1{
		t := int(int(num1[l1]) - int('0') + int(num2[l2]) - int('0') +add)
		if t >=10{
			t -= 10
			add = 1
		}else{
			add = 0
		}
		ret = string(t+'0') + ret
	}
	if l1 < 0{
		l1, num1 = l2,num2
	}
	for ;l1>=0;l1--{
		t := int(int(num1[l1]) - int('0')  + add)
				if t >=10{
			t -= 10
			add = 1
		}else{
			add = 0
		}
		ret = string(t+'0') + ret
	}
	if add == 1{
		ret = "1" + ret
	}
	return ret;
}