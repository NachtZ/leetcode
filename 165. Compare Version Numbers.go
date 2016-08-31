package main

import "fmt"

//import "unicode"
import "strings"
import "strconv"
/***************************************/
//some common data struct and function using in leetcode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

type Point struct {
	X int
	Y int
}

func useLib(){
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strings.Compare("1","2"))
}

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = nums[0]
	ch := make(chan *TreeNode, len(nums))
	ch <- root
	nums = nums[1:]
	for i := 0; i < len(nums); i++ {
		tree := <-ch
		if nums[i] == -1 {
			tree.Left = nil
		} else {
			tree.Left = &TreeNode{
				Val: nums[i],
			}
			ch <- tree.Left
		}
		i++
		if i == len(nums) || nums[i] == -1 {
			tree.Right = nil
		} else {
			tree.Right = &TreeNode{
				Val: nums[i],
			}
			ch <- tree.Right
		}
	}
	return root
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	root := &ListNode{
		Val: nums[0],
	}
	tmp := root
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{
			Val: nums[i],
		}
		tmp = tmp.Next
	}
	return root
}
/**************************************************************/
func compareVersion(version1 string, version2 string) int {
    v1,v2 := strings.Split(version1,"."),strings.Split(version2,".")
	ret := 1
	i :=0
	for i=0;i<len(v1)&& i<len(v2);i++{
		a,_ := strconv.Atoi(v1[i])
		b,_ := strconv.Atoi(v2[i])
		if a>b{
			return 1
		}else if a < b{
			return -1
		}
	}
	if len(v1)<len(v2){
		ret = -1
		v1 = v2
	}
	for ;i<len(v1);i++{
		a,_:= strconv.Atoi(v1[i])
		if a!=0{
			return ret
		}
	}
	return 0
}
func main() {
	fmt.Println(compareVersion("1","1.0"))
}
