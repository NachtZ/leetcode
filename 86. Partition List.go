package main
import "fmt"
  type ListNode struct {
      Val int
      Next *ListNode
  }
 
func partition(head *ListNode, x int) *ListNode {
    var sHead,bHead ListNode
    sHead.Next,bHead.Next = nil,nil
    var sTmp,bTmp *ListNode
    sTmp,bTmp = nil,nil
    for head!= nil{
        if head.Val < x{
            if sHead.Next == nil{
                sHead.Next = head
                sTmp = head
            }else{
                sTmp.Next = head
                sTmp = sTmp.Next
            }
        }else{
            if bHead.Next == nil{
                bHead.Next = head
                bTmp = head
            }else{
                bTmp.Next = head
                bTmp = bTmp.Next
            }
        } 
    }
    if sTmp == nil{
        return bHead.Next
    }
    if bTmp == nil{
        return sHead.Next
    }
    sTmp.Next,bTmp.Next =bHead.Next,nil
    return sHead.Next
    
}
func main(){
    str := []string{"10100","10111","11111","10010"}
    m:= [][]byte{}
    for i:=0;i<len(str);i++{
        tmp :=[]byte{}
        for _,j := range str[i]{
            tmp = append(tmp,byte(j))
        }
        m = append(m,tmp)
    }
    fmt.Println(maximalRectangle(m))
    fmt.Print("Hello")
}
