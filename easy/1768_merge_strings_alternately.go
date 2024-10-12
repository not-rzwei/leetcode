//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=1768 lang=golang
 *
 * [1768] Merge Strings Alternately
 */
package easy

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	// pointers to the start of word1 and word2
	p1, p2 := 0, 0

	// reserve memory for the merged string
	merged := make([]byte, 0, len(word1)+len(word2))

	// loop until array is filled
	for len(merged) < cap(merged) {
		// add letter of word1
		if p1 < len(word1) {
			merged = append(merged, word1[p1])
			p1++
		}

		// add letter of word2
		if p2 < len(word2) {
			merged = append(merged, word2[p2])
			p2++
		}
	}

	return string(merged)
}

// @lc code=end
