package main
import "fmt"
//import "unicode"
//import "strings"
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

func help(board [][]byte,visit [][]bool,x,y int)bool{
    if x <0 || x>= len(board) || y <0 || y>=len(board[0]){
        return false
    }
    if visit[x][y]{
        return board[x][y] != 'O'
    }
    visit[x][y] = true
    flag := true
    if board[x][y] == 'O'{
        board[x][y] = 'X'//sharp means pending
        flag = help(board,visit,x+1,y) && help(board,visit,x-1,y) && help(board,visit,x,y+1) && help(board,visit,x,y-1)
        if flag == false{
            board[x][y] = 'O'
        }
    }
    return flag

}


func solve(board [][]byte)  {
    l1 := len(board)
    if l1 == 0{
        return
    }
    l2 := len(board[0])
    if l2 == 0{
        return
    }
    visit := make([][]bool,l1)
    for i :=0;i<l1;i++{
        visit[i] = make([]bool,l2)
    }
    for i:=0;i<l1;i++{
        for j:=0;j<l2;j++{
            if !visit[i][j]{
                help(board,visit,i,j)
            }
        }
    }
}



func main(){
    strs := []string{"XOOXXXOXOO","XOXXXXXXXX","XXXXOXXXXX","XOXXXOXXXO","OXXXOXOXOX","XXOXXOOXXX","OXXOOXOXXO","OXXXXXOXXX","XOOXXOXXOO","XXXOOXOXXO"}
    input := [][]byte{}
    for _,str := range strs{
        input = append(input,[]byte(str))
    }
    solve(input)
    fmt.Println(input)
    fmt.Print("Hello")
}
