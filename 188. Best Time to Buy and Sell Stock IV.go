package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
)

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
	fmt.Println(math.Abs(1.0))
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
func maxProfit(k int, prices []int) int {
    if k ==0{
		return 0
	}
	if k >= len(prices)/2{
		res :=0
		for i:=1;i<len(prices);i++{
			if prices[i]-prices[i-1]>0{
				res += prices[i]-prices[i-1]
			}
		}
		return res
	}
	buy,sell := make([]int,k),make([]int,k)
	for i:=0;i<k;i++{
		buy[i] = -65535000
	}
	for _,price := range prices{
		if buy[0]< -price{
			buy[0] = -price
		}
		if sell[0]<buy[0]+price{
			sell[0] = buy[0]+price
		}
		for i:=1;i<k;i++{
			if sell[i-1]-price >buy[i]{
				buy[i] = sell[i-1]-price
			}
			if sell[i]< buy[i]+price{
				sell[i] = buy[i]+price
			}
		}
	}
	return sell[k-1]
}
func main() {
	fmt.Println(maxProfit(2,[]int{3,2,6,5,0,3}))
}
