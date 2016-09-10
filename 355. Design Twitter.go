package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"unicode"
	//"math/rand"
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



func max(a,b int)int{
	if a<b{
		return b
	}
	return a
}
func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}
/**************************************************************/
type Twitter struct {
	user map[int][]int//key:user id,value:tweets
	follow map[int]map[int]bool//key follower, value: be followerd
	tweets [][]int//time line, []int[0]user,[]int[1]tweet id
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	var t Twitter
	t.user =make(map[int][]int)
	t.follow = make(map[int]map[int]bool)
	t.tweets = [][]int{}
	return t
}


/** Compose a new tweet. */
func (this *Twitter) Posttweet(userId int, tweetId int)  {
	this.tweets = append(this.tweets,[]int{userId,tweetId})
	_,ok := this.user[userId]
	if ok{
		this.user[userId] = append(this.user[userId],tweetId)
	}else{
		this.user[userId] = []int{tweetId}
	}
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) Getnewsfeed(userId int) []int {
	count := 0
	ret := []int{}
	//this.follow[userId][userId] = true
	
	if len(this.follow[userId]) == 0{
		this.follow[userId] = make(map[int]bool)
		this.follow[userId][userId] = true
	}
	m := this.follow[userId]
	for i:=len(this.tweets)-1;i>=0&&count <10;i--{
		if m[this.tweets[i][0]] == true{
			count ++
			ret = append(ret,this.tweets[i][1])
		}
	}
	return ret
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	if len(this.follow[followerId]) == 0{
		this.follow[followerId] = make(map[int]bool)
		this.follow[followerId][followerId] = true
	}
	this.follow[followerId][followeeId] = true
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if followeeId == followerId{
		return
	}
	if len(this.follow[followerId]) == 0{
		this.follow[followerId] = make(map[int]bool)
		this.follow[followerId][followerId] = true
		return 
	}
	this.follow[followerId][followeeId] = false
}

func main() {
	t := Constructor()
	t.Posttweet(1,5)
	fmt.Println(t.Getnewsfeed(1))
	t.Follow(1,2)
	t.Posttweet(2,6)
	fmt.Println(t.Getnewsfeed(1))
	t.Unfollow(1,2)
	fmt.Println(t.Getnewsfeed(1))
	//fmt.Println([][]int{},0)
}
