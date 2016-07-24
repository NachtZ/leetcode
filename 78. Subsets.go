package main

import "fmt"

//import "math"

func subsets(nums []int) [][]int {

	if len(nums) == 0 {
		return [][]int{}
	}
	max := 1 << uint16(len(nums))
	//max := int(math.Pow(2, float64(len(nums))))
	ret := [][]int{}
	for i := 0; i < max; i++ {
		bit := i
		tmp := []int{}
		for j := 0; j < len(nums); j++ {
			if bit&1 == 1 {
				tmp = append(tmp, nums[j])
			}
			bit >>= 1
		}
		ret = append(ret, tmp)
	}
	return ret
}
func main() {
	//fmt.Println(isNumber("123456"))
	//fmt.Print(1 << 10)
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println("Hello")
}
