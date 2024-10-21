//go:build cheatsheet
// +build cheatsheet

package cheatsheet

/** standard lib heap interface
 type sort.Interface {
	Len() int
	Less(i, j int) bool // min/max heap decision factor
	Swap(i, j int)
 }

 type Interface {
	sort.Interface
	Push(x any) // append new element to array
	Pop() // pop last element from array
 }
*/

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	lastIdx := len(*h) - 1
	item := (*h)[lastIdx]
	*h = (*h)[:lastIdx]
	return item
}
