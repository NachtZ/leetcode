package main

import "fmt"

type MyCalendarTwo struct {
	over [][2]int
	cal  [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		over: [][2]int{},
		cal:  [][2]int{},
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, e := range this.over {
		if e[1] > start && e[0] < end {
			return false
		}
	}
	for _, e := range this.cal {
		if e[1] > start && e[0] < end {
			i, j := start, end
			if start < e[0] {
				i = e[0]
			}
			if end > e[1] {
				j = e[1]
			}
			this.over = append(this.over, [2]int{i, j})
		}
	}
	this.cal = append(this.cal, [2]int{start, end})
	return true
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
