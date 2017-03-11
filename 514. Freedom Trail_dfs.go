package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
	//"math/rand"
	//"sort"
	//"regexp"
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

type Interval struct {
	Start int
	End   int
}

func useLib() {
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strings.Compare("1", "2"))
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**************************************************************/
/*
type IntSlice []int

func (s IntSlice)Less(i,j int)bool{
	return s[i]<s[j]
}
func (s IntSlice)Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}
func (s IntSlice)Len()int{
	return len(s)
}
*/
/**************************************************************/
var res int
var maps map[[2]int]int

/*
	for i := 0; i < len(key); i++ {
		ret += moves[idx][key[i]] //rotate
		ret++                     //spell
		t := (moves[idx][key[i]] + idx) % len(ring)
		if ring[t] != key[i] {
			t = (idx - moves[idx][key[i]] + len(ring)) % len(ring)
		}
		idx = t
	}
*/

func dfs(ring string, key string, idx, i int, left, right []map[byte]int, count []int, ret int) {
	if v, ok := maps[[2]int{idx, i}]; ok {
		if v <= ret {
			return
		}
	}
	maps[[2]int{idx, i}] = ret
	if ret > res {
		return
	}
	if i == len(key) {
		if ret < res {
			res = ret
		}
		return
	}
	if ring[idx] == key[i] {
		if ret+1 < res {
			dfs(ring, key, idx, i+1, left, right, count, ret+1)
		}
		return

	}
	if count[key[i]] == 1 {
		t := (left[idx][key[i]] + idx) % len(ring)
		if ring[t] != key[i] {
			t = (idx - left[idx][key[i]] + len(ring)) % len(ring)
		}
		if ret+left[idx][key[i]]+1 < res {
			dfs(ring, key, t, i+1, left, right, count, ret+left[idx][key[i]]+1)
		}
	} else {
		if _, ok := left[idx][key[i]]; ok && ret+left[idx][key[i]]+1 < res {
			dfs(ring, key, (idx+left[idx][key[i]])%len(ring), i+1, left, right, count, ret+left[idx][key[i]]+1)
		}
		if _, ok := right[idx][key[i]]; ok && ret+right[idx][key[i]]+1 < res {
			dfs(ring, key, (idx-right[idx][key[i]]+len(ring))%len(ring), i+1, left, right, count, ret+right[idx][key[i]]+1)
		}
	}
}

func findRotateSteps(ring string, key string) int {
	moveLeft := make([]map[byte]int, len(ring))
	moveRight := make([]map[byte]int, len(ring))
	maps = make(map[[2]int]int)
	count := make([]int, 256)
	for i := 0; i < len(ring); i++ {
		count[ring[i]]++
	}
	for i := 0; i < len(ring); i++ {
		moveLeft[i] = make(map[byte]int)
		moveRight[i] = make(map[byte]int)
		for j := 0; j < len(ring); j++ {
			v := ring[(j+i)%len(ring)]
			if v == ring[i] {
				moveLeft[i][v] = 0
				continue
			}
			if _, ok := moveLeft[i][v]; !ok {
				moveLeft[i][v] = j
			}
		}
		for j := 0; j < len(ring); j++ {
			v := ring[(i-j+len(ring))%len(ring)]
			if v == ring[i] {
				moveRight[i][v] = 0
				continue
			}
			if _, ok := moveRight[i][v]; !ok {
				moveRight[i][v] = j
			}
			if count[v] == 1 && moveLeft[i][v] > moveRight[i][v] {
				moveLeft[i][v] = moveRight[i][v]
			}
		}
	}
	/*
		for _, m := range moveLeft {
			for k, v := range m {
				fmt.Print(string(k), ":", v, " ")
			}
			fmt.Println()
		}
		fmt.Println("---")
		for _, m := range moveRight {
			for k, v := range m {
				fmt.Print(string(k), ":", v, " ")
			}
			fmt.Println()
		}*/
	/*
		ret, idx := 0, 0
		for i := 0; i < len(key); i++ {
			ret += moves[idx][key[i]] //rotate
			ret++                     //spell
			t := (moves[idx][key[i]] + idx) % len(ring)
			if ring[t] != key[i] {
				t = (idx - moves[idx][key[i]] + len(ring)) % len(ring)
			}
			idx = t
		}*/
	res = 1000000
	dfs(ring, key, 0, 0, moveLeft, moveRight, count, 0)
	return res
}

func main() {
	//fmt.Println(findRotateSteps("godding", "godding"))
	fmt.Println(findRotateSteps("caotmcaataijjxi", "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx"))
}
