package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"unicode"
	"math/rand"
	//"sort"
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
	fmt.Println(unicode.IsDigit('1'))
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

func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}

func max(a,b int)int{
	if a<b{
		return b
	}
	return a
}

/**************************************************************/
type RandomizedCollection struct {
    arr []int 
	del []int
	total int
	m map[int][]int
}


/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
    var r RandomizedCollection
	r.m = make(map[int][]int)
	return r
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	idx := len(this.arr)
    if len(this.del)!=0{
		idx = this.del[0]
		this.del = this.del[1:]
		this.arr[idx] =val 
	}else{
		this.arr = append(this.arr,val)
	}
	_,ok := this.m[val]
	if ok{
		this.m[val] = append(this.m[val],idx)
	}else{
		this.m[val] = []int{idx}
	}
	this.total ++
	return !ok
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	_,ok := this.m[val]
	if ok == false{
		return false
	}
	this.total --
	idx := this.m[val][0]
	if len(this.m[val])==1{
		delete(this.m,val)
	}else{
		this.m[val] = this.m[val][1:]
	}
	this.del = append(this.del,idx)
	this.arr[idx] = -65535
	return true
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
    
	for {
		r := rand.Intn(len(this.arr))
		if this.arr[r] != -65535{
			return this.arr[r]
		}
	}
}


/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
func main() {
	t := Constructor()
	fmt.Println(t.Insert(1))
	fmt.Println(t.Insert(1))
	fmt.Println(t.Remove(1))
	fmt.Println(t.GetRandom())
	//fmt.Println([][]int{},0)
}
