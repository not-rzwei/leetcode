package easy

import (
	. "leetcode/helper"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		tree []interface{}
		want int
	}{
		{
			tree: []interface{}{},
			want: 0,
		},
		{
			tree: []interface{}{1},
			want: 0,
		},
		{
			tree: []interface{}{1, nil, 2},
			want: 1,
		},
		{
			tree: []interface{}{1, 2, 3, 4, 5},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(SliceToString(tt.tree), func(t *testing.T) {
			tree := BuildBTree(tt.tree, 0)
			if got := diameterOfBinaryTree(tree); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
