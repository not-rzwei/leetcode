//go:build iterative || recursive
// +build iterative recursive

package easy

import (
	"fmt"
	. "leetcode/helper"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	tests := []struct {
		p    []interface{}
		q    []interface{}
		want bool
	}{
		{
			p:    []interface{}{},
			q:    []interface{}{},
			want: true,
		},
		{
			p:    []interface{}{1, 2, 3},
			q:    []interface{}{1, 2, 3},
			want: true,
		},
		{
			p:    []interface{}{1, 2},
			q:    []interface{}{1, nil, 3},
			want: false,
		},
		{
			p:    []interface{}{1, 2, 1},
			q:    []interface{}{1, 1, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		title := fmt.Sprintf("P %s Q %s", SliceToString(tt.p), SliceToString(tt.q))
		t.Run(title, func(t *testing.T) {
			p := BuildBTree(tt.p, 0)
			q := BuildBTree(tt.q, 0)
			if got := isSameTree(p, q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
