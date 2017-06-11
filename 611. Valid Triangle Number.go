package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
	//"math/rand"
	//"sort"
	"regexp"
	"sort"
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
	fmt.Println(sort.IsSorted(nil))
	fmt.Println(regexp.QuoteMeta(""))
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

func printTreeNode(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}
	left, right := "nil", "nil"
	if root.Left != nil {
		left = strconv.Itoa(root.Left.Val)
	}
	if root.Right != nil {
		right = strconv.Itoa(root.Right.Val)
	}
	fmt.Println(root.Val, ":", left, right)
	printTreeNode(root.Left)
	printTreeNode(root.Right)
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
func printList(root *ListNode) {
	for root != nil {
		fmt.Print(root.Val, "->")
		root = root.Next
	}
	fmt.Println("nil")
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

type trie struct {
	ch    byte
	child [26]*trie
	flag  bool // true means here located a word.
}

func (this *trie) insertWords(str string) {
	ptr := this
	for idx, s := range str {
		if ptr.child[s-'a'] == nil {
			ptr.child[s-'a'] = new(trie)
			ptr.child[s-'a'].ch = str[idx]
		}
		if idx == len(str)-1 {
			ptr.child[s-'a'].flag = true
		}
		ptr = ptr.child[s-'a']
	}
}

func (this *trie) searchWords(str string) *trie {
	ptr := this
	for idx, s := range str {
		if ptr.child[s-'a'] == nil {
			return nil
		}
		if idx == len(str)-1 {
			if ptr.child[s-'a'].flag == true {
				return ptr.child[s-'a']
			}
			return nil
		}
		ptr = ptr.child[s-'a']

	}
	return nil
}

func buildTrie(words []string) *trie {
	root := &trie{}
	for _, w := range words {
		root.insertWords(w)
	}
	return root
}

/**************************************************************/
type IntSlice []int

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntSlice) Len() int {
	return len(s)
}

type StrSlice []string

func (s StrSlice) Less(i, j int) bool {
	if len(s[i]) != len(s[j]) {
		return len(s[i]) > len(s[j])
	}
	return s[i] < s[j]
}
func (s StrSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StrSlice) Len() int {
	return len(s)
}

/**************************************************************/

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

/*************************************************************/
func triangleNumber(nums []int) int {
	cnt := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		beforek := i + 1
		for j := i + 1; j < len(nums)-1; j++ {
			if beforek < j+1 {
				beforek = j + 1
			}
			cnt += beforek - j - 1
			for k := beforek; k < len(nums); k++ {
				if nums[k] < nums[i]+nums[j] && nums[i] > nums[k]-nums[j] {
					cnt++
				} else {
					beforek = k
					break
				}
			}
		}
	}
	return cnt
}
func main() {
	fmt.Println(triangleNumber([]int{292, 445, 756, 527, 508, 104, 523, 601, 462, 249, 224, 640, 411, 806, 96, 761, 348, 546, 451, 965, 53, 95, 158, 851, 873, 491, 10, 993, 736, 129, 718, 527, 475, 400, 601, 713, 252, 673, 647, 791, 340, 506, 656, 359, 266, 559, 411, 630, 28, 685, 315, 160, 563, 454, 184, 524, 132, 939, 904, 385, 375, 901, 549, 261, 377, 541, 803, 840, 417, 595, 833, 75, 976, 424, 83, 929, 197, 697, 431, 257, 431, 386, 446, 276, 199, 530, 334, 888, 601, 158, 0, 904, 85, 678, 616, 631, 143, 740, 321, 488, 357, 887, 682, 602, 321, 894, 193, 149, 239, 78, 619, 373, 254, 920, 717, 288, 242, 76, 568, 542, 413, 64, 506, 695, 574, 942, 726, 331, 262, 794, 997, 35, 382, 364, 878, 86, 646, 245, 255, 903, 481, 579, 470, 284, 642, 610, 872, 960, 902, 995, 520, 619, 131, 775, 99, 187, 634, 899, 99, 390, 209, 225, 337, 391, 408, 588, 626, 405, 832, 679, 323, 787, 200, 757, 253, 749, 847, 783, 171, 306, 989, 300, 620, 47, 809, 521, 782, 932, 398, 108, 556, 171, 239, 225, 79, 441, 211, 433, 56, 926, 665, 238, 749, 718, 75, 768, 298, 82, 685, 269, 287, 947, 494, 435, 769, 79, 282, 870, 168, 736, 786, 420, 768, 838, 252, 337, 391, 376, 294, 175, 317, 944, 269, 35, 274, 889, 750, 664, 397, 511, 189, 711, 764, 64, 12, 880, 518, 828, 211, 154, 582, 43, 661, 91, 46, 714, 116, 762, 718, 78, 105, 118, 227, 403, 569, 650, 394, 100, 299, 904, 71, 723, 888, 694, 421, 828, 919, 369, 960, 304, 757, 338, 128, 204, 133, 513, 516, 847, 532, 225, 433, 563, 471, 57, 463, 641, 400, 857, 165, 103, 77, 291, 430, 956, 388, 522, 855, 147, 668, 885, 532, 314, 877, 47, 301, 975, 36, 569, 197, 929, 135, 216, 998, 584, 636, 243, 625, 926, 462, 525, 408, 737, 967, 626, 159, 954, 286, 33, 232, 616, 254, 381, 197, 420, 863, 247, 268, 486, 714, 44, 700, 862, 241, 136, 462, 670, 402, 446, 936, 786, 224, 603, 601, 511, 259, 772, 748, 644, 619, 511, 920, 736, 757, 723, 187, 676, 197, 13, 688, 797, 963, 314, 798, 666, 594, 864, 84, 801, 200, 861, 502, 187, 477, 245, 273, 798, 15, 307, 838, 338, 340, 649, 146, 631, 229, 93, 5, 311, 723, 525, 566, 532, 819, 994, 171, 597, 949, 963, 836, 482, 140, 843, 55, 938, 700, 713, 417, 722, 968, 596, 94, 476, 627, 804, 401, 178, 853, 60, 503, 764, 449, 817, 147, 668, 349, 503, 252, 735, 440, 592, 487, 772, 457, 881, 263, 444, 994, 45, 763, 951, 678, 2, 144, 288, 63, 640, 541, 59, 707, 491, 765, 947, 255, 862, 358, 19, 41, 369, 763, 550, 184, 195, 955, 119, 122, 694, 821, 957, 973, 10, 437, 297, 714, 925, 860, 271, 768, 570, 861, 107, 907, 500, 334, 241, 934, 772, 95, 675, 381, 512, 883, 163, 631, 233, 360, 45, 532, 45, 9, 942, 735, 775, 389, 668, 703, 968, 608, 518, 432, 923, 204, 876, 216, 362, 99, 675, 773, 219, 52, 149, 309, 747, 9, 789, 455, 851, 597, 534, 463, 549, 723, 867, 695, 4, 742, 745, 540, 766, 93, 50, 806, 143, 979, 615, 19, 356, 253, 706, 855, 248, 927, 320, 646, 738, 14, 362, 237, 510, 661, 881, 637, 387, 313, 273, 37, 748, 153, 785, 132, 797, 177, 895, 128, 204, 808, 58, 324, 119, 344, 697, 697, 473, 750, 286, 768, 588, 844, 372, 621, 739, 58, 932, 206, 454, 103, 8, 321, 880, 425, 599, 387, 698, 27, 798, 65, 856, 507, 574, 739, 366, 678, 297, 566, 209, 385, 340, 352, 434, 106, 788, 634, 475, 477, 850, 214, 896, 548, 616, 703, 414, 448, 723, 165, 832, 793, 795, 421, 519, 23, 210, 770, 344, 0, 899, 160, 809, 909, 549, 965, 873, 512, 604, 757, 407, 714, 674, 753, 793, 946, 45, 69, 855, 797, 441, 842, 397, 223, 435, 245, 73, 884, 531, 7, 515, 438, 625, 114, 833, 670, 408, 185, 396, 716, 120, 836, 369, 932, 175, 993, 708, 285, 295, 412, 818, 80, 615, 805, 774, 87, 149, 385, 438, 211, 332, 560, 755, 148, 951, 584, 365, 952, 157, 613, 186, 869, 339, 698, 175, 664, 775, 816, 559, 773, 834, 203, 749, 183, 683, 713, 98, 193, 340, 810, 291, 466, 600, 691, 312, 330, 303, 842, 490, 855, 851, 722, 628, 5, 693, 331, 267, 403, 126, 725, 248, 122, 112, 298, 802, 587, 165, 990, 682, 187, 933, 108, 418, 80, 384, 621, 729, 209, 639, 741, 358, 93, 314, 816, 148, 670, 94, 166, 865, 397, 559, 812, 125, 671, 893, 740, 364, 632, 844, 228, 52, 284, 745, 581, 383, 526, 333, 492, 670, 484, 682, 867, 407, 541, 50, 800, 233, 986, 642, 941, 549, 499, 107, 321, 557, 380, 984, 648, 436, 105, 537, 38, 882, 329, 901, 20, 757, 157, 534, 472, 214, 290, 88, 748, 470, 437, 903, 666, 255, 852, 181, 401, 723, 169, 703, 803, 13, 987, 864, 576, 724, 344, 588, 765, 456, 771, 852, 61, 894, 209, 515, 389, 980, 131, 227, 704, 986, 363, 138, 173, 847, 118, 388, 941, 175, 654, 295, 213, 166, 132, 320, 567, 513, 989, 481, 666, 873, 474, 921, 860, 865, 157, 174, 828, 920, 374, 736, 605, 652, 671, 638, 306, 390, 864, 116, 408, 556, 76, 961, 282, 903, 159, 351, 1, 558, 699, 169, 767, 679, 252, 696, 104, 759, 91, 585, 205, 956, 314, 19, 303, 566, 530, 395, 783, 502, 553, 520, 816, 584, 374, 235, 667, 541, 837, 63, 747, 226, 103, 611, 84, 214, 853, 589, 223, 479, 778, 766, 491, 554, 290, 649, 281, 184, 855, 760, 18, 174, 372, 598, 205, 135, 58, 357, 0, 163, 376, 116}))
}
