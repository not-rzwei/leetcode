//go:build backward || forward
// +build backward forward

/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

package medium

import (
	"leetcode/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dailyTemperatures(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{
			args: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			args: []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70},
			want: []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(helper.SliceToString(tt.args), func(t *testing.T) {
			got := dailyTemperatures(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
