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

	// track the letter count using byte array
	// 26 is the number of english letters in lowercase
	letterCount := [26]byte{}

	// loop through the string
	// whichever is fine
	for i := 0; i < len(s); i++ {
		// increase the counter for s and decrease it for t
		// substract the index with -a to offset the ASCII byte
		// a is start from 97 and we offset with -97 to make it starts from 0
		letterCount[s[i]-'a']++
		letterCount[t[i]-'a']--
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
