package main
import "fmt"
import "sort"

 type Interval struct {
 	   Start int
 	   End   int
  }

type IL []Interval
func (list IL) Len()int{
    return len(list)
}
func (list IL) Less(i,j int)bool{
    if  list[i].Start == list[j].Start {
        return list[i].End < list[j].End
    }
    return list[i].Start < list[j].Start
}



func (list IL) Swap(i,j int){
    tmp := list[i]
    list[i] = list[j]
    list[j] = tmp
}


func merge(intervals []Interval) []Interval {
    if len(intervals) <=1{
        return intervals
    }
    var il IL
    ret := []Interval{}
    il = intervals

    sort.Sort(il)
    tmp := il[0]
    for i:=1;i<len(il);i++{
        if il[i].End < tmp.Start || il[i].Start > tmp.End{
     //       tmp1 := tmp           
            ret =append(ret,tmp)
            tmp = il[i]
        }else{
            if tmp.Start > il[i].Start{
                tmp.Start = il[i].Start
            }
            if tmp.End < il[i].End{
                tmp.End = il[i].End
            }
        }
    }
    ret = append(ret,tmp)
    return ret

}

func insert(intervals []Interval, newInterval Interval) []Interval {
    intervals = append(intervals,newInterval)
    return merge(intervals)
}

func main(){
    var tmp Interval
    arg := []Interval{}
    tmp.Start = 1
    tmp.End = 4
    arg = append(arg,tmp)
    arg = append(arg,tmp)
    merge(arg)
    fmt.Print("Hello")
}
