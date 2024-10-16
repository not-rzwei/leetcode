//go:build iterative || recursive
// +build iterative recursive

package easy

import (
	"fmt"
	. "leetcode/helper"
	"testing"
)

func Test_isSubtree(t *testing.T) {
	tests := []struct {
		root    []interface{}
		subroot []interface{}
		want    bool
	}{
		{
			root:    []interface{}{3, 4, 5, 1, 2},
			subroot: []interface{}{4, 1, 2},
			want:    true,
		},
		{
			root:    []interface{}{3, 4, 5, 1, 2, nil, nil, nil, nil, 0},
			subroot: []interface{}{4, 1, 2},
			want:    false,
		},
		{
			root:    []interface{}{1},
			subroot: []interface{}{0},
			want:    false,
		},
		{
			root:    []interface{}{1},
			subroot: []interface{}{1},
			want:    true,
		},
		{
			root:    []interface{}{1, 1},
			subroot: []interface{}{1},
			want:    true,
		},
	}
	for _, tt := range tests {
		title := fmt.Sprintf("ROOT %s SUBROOT %s", SliceToString(tt.root), SliceToString(tt.subroot))
		t.Run(title, func(t *testing.T) {
			p := BuildBTree(tt.root, 0)
			q := BuildBTree(tt.subroot, 0)
			if got := isSubtree(p, q); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
