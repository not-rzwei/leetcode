//go:build heap
// +build heap

package easy

import (
	"leetcode/helper"
	"testing"
)

func Test_lastStoneWeight(t *testing.T) {
	tests := []struct {
		stones []int
		want   int
	}{
		{
			stones: []int{2, 7, 4, 1, 8, 1},
			want:   1,
		},
		{
			stones: []int{1},
			want:   1,
		},
		{
			stones: []int{1, 1},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(helper.SliceToString(tt.stones), func(t *testing.T) {
			if got := lastStoneWeight(tt.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
