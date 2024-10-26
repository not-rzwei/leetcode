//go:build sort
// +build sort

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

package medium

import (
	"slices"
)

// @lc code=start
func groupAnagrams(strs []string) (result [][]string) {
	groups := make(map[string][]string)

	// use sorted string as grouping key
	// and assign strs to respective group
	for _, str := range strs {
		key := sortString(str)
		groups[key] = append(groups[key], str)
	}

	// add each group to result
	for _, group := range groups {
		result = append(result, group)
	}

	return
}

// sort the slice of string character
func sortString(s string) string {
	key := []byte(s)
	slices.Sort(key)
	return string(key)
}

// @lc code=end
