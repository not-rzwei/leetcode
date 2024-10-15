//go:build recursive || iterative
// +build recursive iterative

package easy

import (
	. "leetcode/helper"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		tree []interface{}
		want bool
	}{
		{
			tree: []interface{}{},
			want: true,
		},
		{
			tree: []interface{}{1},
			want: true,
		},
		{
			tree: []interface{}{1, 2, 3, 4, 5},
			want: true,
		},
		{
			tree: []interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(SliceToString(tt.tree), func(t *testing.T) {
			tree := BuildBTree(tt.tree, 0)
			if got := isBalanced(tree); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
