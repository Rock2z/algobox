package _347

import (
	"container/heap"
	"sort"
)

type myHeap struct {
	s [][]int
}

func (h *myHeap) Len() int {
	return len(h.s)
}

func (h *myHeap) Less(i, j int) bool {
	return h.s[i][1] < h.s[j][1]
}

func (h *myHeap) Swap(i, j int) {
	h.s[i], h.s[j] = h.s[j], h.s[i]
}

func (h *myHeap) Push(x interface{}) {
	h.s = append(h.s, x.([]int))
}

func (h *myHeap) Pop() interface{} {
	ans := h.s[len(h.s)-1]
	h.s = h.s[:len(h.s)-1]
	return ans
}

func topKFrequent(nums []int, k int) []int {
	h := &myHeap{s: make([][]int, 0)}
	heap.Init(h)
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for key, val := range m {
		heap.Push(h, []int{key, val})
		// use min heap, so every time heap length bigger than k, need to pop out extra elements
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(h).([]int)[0])
	}
	return ret
}

func topKFrequent0(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	ret := make([]int, 0, len(m))
	for k, _ := range m {
		ret = append(ret, k)
	}
	sort.Slice(ret, func(i, j int) bool {
		return m[ret[i]] > m[ret[j]]
	})
	return ret[:k]
}
