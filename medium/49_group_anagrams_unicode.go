//go:build unicode
// +build unicode

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

package medium

// @lc code=start
func groupAnagrams(strs []string) (result [][]string) {
	groups := make(map[UnicodeFrequency][]string)

	// use unicode frequency as grouping key
	// and assign strs to respective group
	for _, str := range strs {
		key := stringToUnicodeFrequency(str)
		groups[key] = append(groups[key], str)
	}

	// add each group to result
	for _, group := range groups {
		result = append(result, group)
	}

	return
}

// shorthand, 26 is the number of lowercase alphabet
type UnicodeFrequency [26]int

// increase chracter unicode code point for each character in string
func stringToUnicodeFrequency(s string) (u UnicodeFrequency) {
	for _, c := range s {
		u[c-'a']++ // -a to start the code point with zero index
	}

	return
}

// @lc code=end
