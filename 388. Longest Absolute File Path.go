package main
import "fmt"
//import "unicode"
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

type node struct{
    path string
    next []*node
}

func lengthLongestPath(input string) int {
    //input = trim(input)
    //input = strings.Replace(input,"\t    ","\t\t",-1)
    input = strings.Replace(input,"\n    ","\n\t",-1)

    paths := strings.Split(input,"\n")
    maxLength := 0
    if strings.Count(paths[0],".") != 0 && strings.Index(paths[0],".") != len(paths[0])-1{
            maxLength = len(paths[0])
    }
    maxPath := ""
    root := &node{
        path:"",
    }
    tmp := root
    for i:=0;i<len(paths);i++{
        tmp = root
        tmplength :=0
        tmppath := ""
        tnode := &node{
            path:"\\"+ strings.TrimLeft(paths[i],"\t"),
        }
        count := strings.Count(paths[i],"\t")
        for count >0{
            count --
                tmppath = tmppath + tmp.path
                tmplength += len(tmp.path)
                tmp = tmp.next[len(tmp.next)-1]
            //}
        }
           // tmppath = tmppath +paths[i]
           // tmplength += len(paths[i])

          //  tmp.next = append(tmp.next,tnode)
        
   // }
            tmppath = tmppath + tmp.path + tnode.path
            tmplength += len(tnode.path) + len(tmp.path)
            if tmplength > maxLength{
                if strings.Count(paths[i],".") != 0 && strings.Index(paths[i],".") != len(paths[i])-1{
                     maxLength = tmplength
                     maxPath = tmppath
                }
               
          //      maxPath = tmppath
            }
            tmp.next = append(tmp.next,tnode)
    }
    fmt.Println(maxLength,maxPath)
    if maxLength >1{
        return maxLength-1
    }else{
        return 0
    }
}
func main(){
    fmt.Println(lengthLongestPath("a"))
    fmt.Print("Hello")
}
