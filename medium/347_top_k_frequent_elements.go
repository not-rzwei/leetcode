//go:build maxheap
// +build maxheap

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

package medium

import "container/heap"

// @lc code=start
type ElementFreq struct {
	Element   int
	Frequency int
}
type MaxHeap []ElementFreq

// Implement heap interface
// Frequency is the priority
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Frequency > h[j].Frequency }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(ElementFreq))
}
func (h *MaxHeap) Pop() any {
	lastIndex := len(*h) - 1
	item := (*h)[lastIndex]
	*h = (*h)[:lastIndex]
	return item
}

func topKFrequent(nums []int, k int) (result []int) {
	freqMap := make(map[int]int)

	// count number frequency
	for _, num := range nums {
		freqMap[num]++
	}

	// heapify the number with its frequency
	freqHeap := MaxHeap{}
	for num, freq := range freqMap {
		heap.Push(&freqHeap, ElementFreq{
			Element:   num,
			Frequency: freq,
		})
	}

	// return only total of k largest element from the heap
	for len(result) < k {
		item := heap.Pop(&freqHeap).(ElementFreq)
		result = append(result, item.Element)
	}

	return
}

// @lc code=end
