package main

import (
	"fmt"
	"math"
	//"sort"
	"strconv"
	"strings"
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

func useLib() {
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strings.Compare("1", "2"))
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
var walkx = [4]int{1, -1, 0, 0}
var walky = [4]int{0, 0, 1, -1}

func help(m, f [][]int, x, y int, flag int) {
	f[x][y] |= flag
	for i := 0; i < 4; i++ {
		_x, _y := x+walkx[i], y+walky[i]
		if _x >= 0 && _x < len(m) && _y >= 0 && _y < len(m[0]) && m[x][y] <= m[_x][_y] && f[_x][_y]&flag == 0 {
			help(m, f, _x, _y, flag)
		}
	}
}

func pacificAtlantic(matrix [][]int) [][]int {
	ret := [][]int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ret
	}
	m, n := len(matrix), len(matrix[0])
	flag := make([][]int, m)
	for i := 0; i < m; i++ {
		flag[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		flag[i][0], flag[i][n-1] = flag[i][0]|1, flag[i][n-1]|2
	}
	for i := 0; i < n; i++ {
		flag[0][i], flag[m-1][i] = flag[0][i]|1, flag[m-1][i]|2
	}
	for i := 0; i < m; i++ {
		help(matrix, flag, i, 0, 1)
		help(matrix, flag, i, n-1, 2)
	}
	for i := 0; i < n; i++ {
		help(matrix, flag, 0, i, 1)
		help(matrix, flag, m-1, i, 2)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if flag[i][j] == 3 {
				ret = append(ret, []int{i, j})
			}
		}
	}
	return ret

}
func main() {
	a := [][]int{}
	/*
		a = append(a, []int{1, 2, 2, 3, 5})
		a = append(a, []int{3, 2, 3, 4, 4})
		a = append(a, []int{2, 4, 5, 3, 1})
		a = append(a, []int{6, 7, 1, 4, 5})
		a = append(a, []int{5, 1, 1, 2, 4})
	*/
	a = append(a, []int{3, 3, 3, 3, 3, 3})
	a = append(a, []int{3, 0, 3, 3, 0, 3})
	a = append(a, []int{3, 3, 3, 3, 3, 3})
	fmt.Println(pacificAtlantic(a))
}
