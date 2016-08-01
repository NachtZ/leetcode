package main
import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
    for t:=m+n-1;t>=0;t--{
        if m == 0 || n == 0{
            if m == 0{
                nums1[t] = nums2[n-1]
                n--
            }else{
                nums1[t] = nums1[m-1]
                m--
            }
            continue
        }
        if nums1[m-1]>nums2[n-1]{
            nums1[t] = nums1[m-1]
            m--
        }else{
            nums1[t] = nums2[n-1]
            n--
        }

    }
}
func main(){
    b := []int{1,2,3,5,7,0,0,0,0,0,0,0}
    merge(b,5,[]int{4,6,7,8,9},5)
    fmt.Println(b)
    fmt.Print("Hello")
}
