package main

import (
	"fmt"
)

type MyCalendar struct {
	order [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		order: [][2]int{},
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, t := range this.order {
		left, right := 0, 0
		if t[0] > start {
			left = t[0]
		} else {
			left = start
		}
		if t[1] > end {
			right = end
		} else {
			right = t[1]
		}
		if left < right {
			return false
		}
	}
	this.order = append(this.order, [2]int{start, end})
	return true
}

func main() {
	t := Constructor()
	fmt.Println(t.Book(10, 20))
	fmt.Println(t.Book(15, 25))
	fmt.Println(t.Book(20, 30))
}
