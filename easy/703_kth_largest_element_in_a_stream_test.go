//go:build heap
// +build heap

package easy

import (
	"leetcode/helper"
	"testing"
)

func TestKthLargest(t *testing.T) {
	tests := []struct {
		k             int
		initialScores []int
		newScores     []int
		want          []int
	}{
		{
			k:             3,
			initialScores: []int{4, 5, 8, 2},
			newScores:     []int{3, 5, 10, 9, 4},
			want:          []int{4, 5, 5, 8, 8},
		},
		{
			k:             2,
			initialScores: []int{4, 5, 8, 2},
			newScores:     []int{3, 1, 10},
			want:          []int{5, 5, 8},
		},
	}
	for _, tt := range tests {
		name := helper.SliceToString(tt.initialScores) + helper.SliceToString(tt.newScores)
		t.Run(name, func(t *testing.T) {
			h := Constructor(tt.k, tt.initialScores)

			for i := 0; i < len(tt.newScores); i++ {
				got := h.Add(tt.newScores[i])
				want := tt.want[i]

				if got != want {
					t.Errorf("want = %v, got = %v", want, got)
				}
			}
		})
	}
}
