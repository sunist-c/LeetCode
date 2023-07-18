package minimum_interval_to_include_each_query

import (
	"container/heap"
	"sort"
)

type Interval struct {
	length int
	end    int
}

type MinHeap struct {
	length int
	data   []Interval
}

func (h *MinHeap) Len() int {
	return h.length
}

func (h *MinHeap) Less(i, j int) bool {
	return h.data[i].length < h.data[j].length || (h.data[i].length == h.data[j].length && h.data[i].end < h.data[j].end)
}

func (h *MinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) Push(x interface{}) {
	h.data = append(h.data, x.(Interval))
	h.length++
}

func (h *MinHeap) Pop() interface{} {
	h.length--
	result := h.data[h.length]
	h.data = h.data[:h.length]
	return result
}

type Query struct {
	idx   int
	value int
}

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	sortedQueries := make([]Query, len(queries))
	for i, q := range queries {
		sortedQueries[i] = Query{i, q}
	}
	sort.Slice(sortedQueries, func(i, j int) bool { return sortedQueries[i].value < sortedQueries[j].value })

	intervalsHeap := &MinHeap{}
	heap.Init(intervalsHeap)
	results := make([]int, len(queries))
	for i := range results {
		results[i] = -1
	}

	i := 0
	for _, query := range sortedQueries {
		for i < len(intervals) && intervals[i][0] <= query.value {
			interval := intervals[i]
			intervalLength := interval[1] - interval[0] + 1
			heap.Push(intervalsHeap, Interval{intervalLength, interval[1]})
			i++
		}

		for intervalsHeap.Len() > 0 && intervalsHeap.data[0].end < query.value {
			heap.Pop(intervalsHeap)
		}

		if intervalsHeap.Len() > 0 {
			results[query.idx] = intervalsHeap.data[0].length
		}
	}

	return results
}
