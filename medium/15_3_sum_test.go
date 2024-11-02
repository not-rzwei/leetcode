/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

package medium

import (
	. "leetcode/helper"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum(t *testing.T) {
	tests := []struct {
		nums       []int
		wantResult [][]int
	}{
		{
			nums: []int{0, 0, 0},
			wantResult: [][]int{
				{0, 0, 0},
			},
		},
		{
			nums: []int{0, 0, 0, 0},
			wantResult: [][]int{
				{0, 0, 0},
			},
		},
		{
			nums:       []int{0, 1, 1},
			wantResult: [][]int{},
		},
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
			wantResult: [][]int{
				{-1, 2, -1},
				{0, 1, -1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(SliceToString(tt.nums), func(t *testing.T) {
			got := threeSum(tt.nums)

			for _, s := range tt.wantResult {
				slices.Sort(s)
			}
			for _, s := range got {
				slices.Sort(s)
			}

			assert.ElementsMatch(t, threeSum(tt.nums), tt.wantResult)
		})
	}
}
