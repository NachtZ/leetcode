package main

import "fmt"
import "sort"

type MyCalendarThree struct {
	m map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		m: make(map[int]int),
	}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	this.m[start], this.m[end] = this.m[start]+1, this.m[end]-1
	a, ret := 0, 0
	tmp := []int{}
	for k, _ := range this.m {
		tmp = append(tmp, k)
	}
	sort.Ints(tmp)
	for _, k := range tmp {
		a += this.m[k]
		if ret < a {
			ret = a
		}
	}
	return ret
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Book(10, 20))
	fmt.Println(obj.Book(50, 60))
	fmt.Println(obj.Book(10, 40))
	fmt.Println(obj.Book(5, 15))
	fmt.Println(obj.Book(5, 10))
	fmt.Println(obj.Book(25, 55))
}
