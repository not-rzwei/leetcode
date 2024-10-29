//go:build twopointers || ignore || twopointers_teleport
// +build twopointers ignore twopointers_teleport

package medium

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "pwwkew",
			want: 3,
		},
		{
			s:    "dvdf",
			want: 3,
		},
		{
			s:    "ckilbkd",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
