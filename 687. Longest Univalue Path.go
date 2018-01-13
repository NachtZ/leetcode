package main

import (
	"fmt"

	//"math/rand"
	//"sort"
)

var res int
//helper return node but not path. path == node +1
func helper(root *TreeNode) int{
	if root == nil {
		return 0
	}
	left,right := helper(root.Left),helper(root.Right)
	if root.Left != nil && root.Right !=nil && root.Left.Val == root.Right.Val && root.Left.Val == root.Val {
		a := left +right + 1
		if a > res{
			res = a
		}
	}
	ret := 0
	if root.Left != nil && root.Left.Val == root.Val{
		ret = left +1
	}else{
		ret = 1
	}
	if root.Right != nil && root.Right.Val == root.Val{
		if ret < right + 1{
			ret = right +1
		}
	}
	if res < ret {
		res = ret
	}
	return ret
}


func longestUnivaluePath(root *TreeNode) int {
	res = 0
	helper(root)
	if res == 0{
		return 0
	}
	return res -1
}
func main() {
	fmt.Println(longestUnivaluePath(nil))
}
