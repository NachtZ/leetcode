/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
type SummaryRanges struct {
    nums []int
}


/** Initialize your data structure here. */
func Constructor() SummaryRanges {
    var s SummaryRanges
    return s
}


func (this *SummaryRanges) Addnum(val int)  {
    if len(this.nums) == 0{
        this.nums = []int{val}
        return
    }
    idx := 0
    for ;idx<len(this.nums);idx++{
        if this.nums[idx] == val{
            return
        }
        if this.nums[idx]>val{
            break
        }
    }
    if idx == len(this.nums){
        this.nums = append(this.nums,val)
    }else if idx == 0{
        this.nums = append([]int{val},this.nums...)
    }else{
        t := append([]int{val},this.nums[idx:]...)
        this.nums = append(this.nums[:idx],t...)
    }
}


func (this *SummaryRanges) Getintervals() []Interval {
    ret := []Interval{}
    if len(this.nums) == 0{
        return ret
    }
    start := 0
    for i:=1;i<len(this.nums);i++{
        if this.nums[i] != this.nums[i-1]+1{
            var t Interval
            t.Start = this.nums[start]
            t.End = this.nums[i-1]
            ret = append(ret,t)
            start = i
        }
    }
    var t Interval
    t.Start = this.nums[start]
    t.End = this.nums[len(this.nums)-1]
    ret = append(ret,t)
    return ret
}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Addnum(val);
 * param_2 := obj.Getintervals();
 */