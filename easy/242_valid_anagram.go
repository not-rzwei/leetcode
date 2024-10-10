//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

package easy

// @lc code=start
func isAnagram(s string, t string) bool {
	// not anagram if they have different length
	if len(s) != len(t) {
		return false
	}

	// track the letter count using map
	letterCount := make(map[byte]int, len(s))

	// loop through the string
	// whichever is fine
	for i := 0; i < len(s); i++ {
		// increase the counter for s and decrease it for t
		letterCount[s[i]]++
		letterCount[t[i]]--
	}

	// loop through the counter
	// if we have non 0 count, strings are not valid anagram
	// because they have extra letter
	for _, count := range letterCount {
		if count != 0 {
			return false
		}
	}

	return true
}

// @lc code=end
