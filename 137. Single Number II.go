func singleNumber(nums []int) int {
    t := len(nums)*8
    bitCount := make([]int,t)
    for i:=0;i<len(nums);i++{
        for j:=0;j<t;j++{
            if 0 != nums[i]&(1<<uint(j)){
                nums[j] +=1 
            }

        }
    }
    single := 0
    for j:=0;j<t;j++{
        if bitCount[j]%3 == 1{
            single += 1<< uint(j)
        }
    }
    return single
}