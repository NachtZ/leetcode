 func help(root * TreeNode, sum,ori int) int{
	if root == nil{
		return 0
	}
	count := 0
    if root.Val == sum{
		count ++
	}
	a,b,c := sum-root.Val, ori-root.Val,ori
	if a<b && a <c{
		if b <c{
			a,b,c = c,b,a
		}else{
			a,b,c = b,c,a
		}
	}else{
		if a > b && b > c{
			if b > c{
				a,b,c = a,b,c
			}else{
				a,b,c = a,c,b
			}
		}else{
			if b >c {
				a,b,c = b,a,c
			}else{
				a,b,c = c,a,b
			}
		}
	}
	fmt.Println(a,b,c)
	if root.Left!=nil{
		count += help(root.Left,a,ori)
		
		if b!=a{
			count += help(root.Left,b,ori)
			//fmt.Print(" ",b)
		}
		if c!=b{
			count += help(root.Left,c,ori)
			//fmt.Print(" ",c)
		}
	}
	if root.Right!=nil{
		count += help(root.Right,a,ori)
		//fmt.Print(" ",a)
		if b!=a{
			count += help(root.Right,b,ori)
			//fmt.Print(" ",b)
		}
		if c!=b{
			count += help(root.Right,c,ori)
			//fmt.Print(" ",c)
		}
	}
	//fmt.Println()
	return count
 }

func pathSum(root *TreeNode, sum int) int {
	return help(root,sum,sum)
}

```
var res int
func go2(root *TreeNode, need int){
	if root == nil{
		return
	}
	if need == 0{
		res ++
	}
	if root.Left != nil{
		go2(root.Left,need-root.Left.Val)
	}
	if root.Right != nil{
		go2(root.Right,need-root.Right.Val)
	}
}
func go1(root * TreeNode, need int){
	if root == nil{
		return
	}
	go2(root,need-root.Val)
	go1(root.Left,need)
	go1(root.Right,need)
}
func pathSum(root *TreeNode, sum int) int {
	res = 0
	go1(root, sum)
	return res
}
```