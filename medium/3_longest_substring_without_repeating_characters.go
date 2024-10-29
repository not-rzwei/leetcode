//go:build twopointers
// +build twopointers

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

package medium

// @lc code=start
// Create a set from map
// Set has function to add, remove rune and check duplicate
type CharSet map[rune]struct{}

func (s *CharSet) Add(c rune) {
	(*s)[c] = struct{}{}
}

func (s *CharSet) Contains(c rune) bool {
	_, exists := (*s)[c]
	return exists
}

func (s *CharSet) Remove(c rune) {
	delete(*s, c)
}

func lengthOfLongestSubstring(s string) int {
	left := 0      // pointer refering to the start of window
	maxLength := 0 // known max length from the window

	// set to track the character inside window
	characterInWindow := make(CharSet)

	// pointer refering to the end of window
	// loop over the string using rune and right pointer
	for right, char := range s {
		// before we can move the right pointer
		// we need to know if the character in right pointer is dupe or not
		// this is to ensure that our right movement is valid
		// if it's already in the window
		// remove that character by removing previous character before it
		// and increase the left pointer (reduce the window)
		for characterInWindow.Contains(char) {
			// the character in right window will be removed by chracter in the left
			// this is to satisfy the character in the window must be unique constraint
			characterInWindow.Remove(rune(s[left]))
			left++
		}

		// add the character to window
		// and then calculate the max window size
		characterInWindow.Add(char)
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

// @lc code=end
