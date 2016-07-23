package main
import "fmt"
func swap(nums []int,i,j int){
    tmp := nums[i]
    nums[i] = nums[j]
    nums[j] = tmp
}
func sortColors(nums []int)  {
    l,i,r := 0,0,len(nums)-1
    for i<=r{
        if nums[i] == 0{
            swap(nums,i,l)
            i++
            l++
        }else {
            if (nums[i] == 1){
                i++
            }else{
                swap(nums,i,r)
                r--
            }

        }
    }
}
func main(){
    fmt.Println(searchMatrix([][]int{{1,1},{2,2}},1))
    fmt.Print("Hello")
}
