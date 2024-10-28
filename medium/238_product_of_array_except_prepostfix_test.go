//go:build prepostfix || prepostfix_inplace
// +build prepostfix prepostfix_inplace

package medium

import (
	"leetcode/helper"
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3},
			want: []int{6, 3, 2},
		},
		{
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(helper.SliceToString(tt.nums), func(t *testing.T) {
			if gotResult := productExceptSelf(tt.nums); !reflect.DeepEqual(gotResult, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}
