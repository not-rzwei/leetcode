//go:build bruteforce || map
// +build bruteforce map

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package easy

import (
	. "leetcode/helper"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 10,
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(SliceToString(tt.nums), func(t *testing.T) {
			if got := twoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
