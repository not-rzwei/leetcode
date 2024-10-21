//go:build cheatsheet
// +build cheatsheet

package cheatsheet_test

import (
	"container/heap"
	"leetcode/helper"
	"testing"
)

func TestMinHeap(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{5, 3, 7, 1, 6},
			want: 1,
		},
		{
			nums: []int{10, -10, -20, 0, 1},
			want: -20,
		},
	}

	for _, tt := range tests {
		t.Run(helper.SliceToString(tt.nums), func(t *testing.T) {
			h := &MinHeap{}

			for _, num := range tt.nums {
				heap.Push(h, num)
			}

			got := heap.Pop(h)
			if got != tt.want {
				t.Errorf("want = %v, got = %v", tt.want, got)
			}
		})
	}
}
