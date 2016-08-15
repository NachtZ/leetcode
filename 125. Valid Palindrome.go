package main
import "fmt"
import "unicode"
import "strings"
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
  type ListNode struct {
      Val int
     Next *ListNode
  }

func buildTree(nums []int)*TreeNode{
    if len(nums) == 0{
        return nil
    }
    root := new(TreeNode)
    root.Val = nums[0]
    ch := make(chan *TreeNode,len(nums))
    ch <- root
    nums =nums[1:]
    for i:=0;i<len(nums);i++{
        tree := <- ch
        if nums[i] == -1{
            tree.Left = nil
        }else{
            tree.Left = &TreeNode{
                Val:nums[i],
            }
            ch <-tree.Left
        }
        i++
        if nums[i] == -1{
            tree.Right = nil
        }else{
            tree.Right = &TreeNode{
                Val:nums[i],
            }
            ch <-tree.Right
        }
    }
    return root
}

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    for i,j := 0,len(s)-1;i<j;{
        if !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])){
            i++
            continue
        }
        if !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])){
            j--
            continue
        }
        if s[i] != s[j]{
            return false
        }
        i++
        j--
    }
    return true
}
func main(){
    fmt.Println(isPalindrome("asdfg"))
    fmt.Print("Hello")
}
