//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

package easy

// @lc code=start
func isSubsequence(s string, t string) bool {
	// pointers refer to the start of s and t
	p1, p2 := 0, 0

	// loop through t with p2 pointer
	for p2 < len(t) {
		// if letter in s match with t
		// increase p1 to check next letter
		if p1 < len(s) && s[p1] == t[p2] {
			p1++
		}

		// move to next p2 letter for checking
		p2++
	}

	// p1 will have the length of s if it's subsequence
	return p1 == len(s)
}

// @lc code=end
