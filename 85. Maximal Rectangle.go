package main
import "fmt"
func maximalRectangle(matrix [][]byte) int {
    m := len(matrix)
    if m ==0{
        return 0
    }
    n := len(matrix[0])
    if n ==0{
        return 0
    }
    result :=0
    height,left,right := make([]int,n),make([]int,n),make([]int,n)
    for i :=0;i<n;i++{
        right[i] = n -1
    }
    for i:=0;i<m;i++{
        LeftNow,RightNow := 0,n-1
        for j:=0;j<n;j++{
            if matrix[i][j] == '1'{
                height[j]++
            }else{
                height[j] = 0
            }
        }
    
         for j:=0;j<n;j++{
            if matrix[i][j] == '1'{
                if LeftNow > left[j]{
                    left[j] = LeftNow
            }
            }else{
                left[j] = 0
                LeftNow = j+1
            }
        }
    
        for j:=n-1;j>=0;j--{
            if matrix[i][j] == '1'{
                if RightNow < right[j]{
                    right[j] = RightNow
                }
            }else{
                right[j] = n-1
                RightNow = j-1
            }
        }
        for j :=0;j<n;j++{
            tmp := (right[j] -left[j]+1) * height[j]
            if result < tmp{
                result = tmp
            }
        }
    }     
    return result
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
