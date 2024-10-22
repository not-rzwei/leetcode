//go:build heap
// +build heap

/*
 * @lc app=leetcode id=1046 lang=golang
 *
 * [1046] Last Stone Weight
 */

package easy

import "container/heap"

// @lc code=start

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // max heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() any {
	lastIndex := len(*h) - 1
	pop := (*h)[lastIndex]
	*h = (*h)[:lastIndex]
	return pop
}

func lastStoneWeight(stones []int) int {
	// kinda sorted for at least 2 first elements
	sortedStones := MaxHeap(stones)

	// loop until we get only 1 stone from stone breakings
	for len(sortedStones) > 1 {
		// heapify the stones
		heap.Init(&sortedStones)

		// pop two biggest stones
		y := heap.Pop(&sortedStones).(int)
		x := heap.Pop(&sortedStones).(int)

		// add the result to heap
		heap.Push(&sortedStones, y-x)
	}

	// no need to handle empty stone
	// because we loop until only 1 stone left
	// so it's guaranteed
	return sortedStones[0]
}

// @lc code=end
