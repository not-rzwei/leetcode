//go:build heap
// +build heap

/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 */

package easy

import "container/heap"

// @lc code=start

type MinHeap []int

// heap implementation
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// push element to array
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// pop last element from array
func (h *MinHeap) Pop() any {
	lastIndex := len(*h) - 1
	item := (*h)[lastIndex]
	*h = (*h)[:lastIndex]
	return item
}

type KthLargest struct {
	Limit  int
	Scores MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	scores := MinHeap{}

	// add score to the score list (heap)
	for _, score := range nums {
		// heap push the score to scores
		// automatically put smallest score to the top of heap
		heap.Push(&scores, score)

		// limit the scores to only k element
		// the smallest one will be removed
		if len(scores) > k {
			heap.Pop(&scores)
		}
	}

	return KthLargest{
		Limit:  k,
		Scores: scores,
	}
}

func (this *KthLargest) Add(val int) int {
	// add score to the heap
	// and limit the heap size
	heap.Push(&this.Scores, val)

	if len(this.Scores) > this.Limit {
		heap.Pop(&this.Scores)
	}

	return this.Scores[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end
