package main

import (
	"container/heap"
	"fmt"
	"time"
)

type WordFrequence struct {
	Word string
	Fre  int
}

type WfHeap []WordFrequence

func (h WfHeap) Len() int {
	return len(h)
}

func (h WfHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h WfHeap) Less(i, j int) bool {
	if h[i].Fre != h[j].Fre {
		return h[i].Fre > h[j].Fre
	}
	return h[i].Word < h[j].Word
}

func (h *WfHeap) Push(x interface{}) {
	*h = append(*h, x.(WordFrequence))
}

func (h *WfHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	h := WfHeap{}
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	for key, val := range m {
		h = append(h, WordFrequence{key, val})
	}
	ptr := &h
	heap.Init(ptr)
	ret := []string{}
	for ; k > 0 && ptr.Len() > 0; k-- {
		ret = append(ret, heap.Pop(ptr).(WordFrequence).Word)
	}
	return ret
}
func main() {
	start := time.Now()
	res := topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4)
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
