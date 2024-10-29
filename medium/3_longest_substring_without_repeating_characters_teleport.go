//go:build twopointers_teleport
// +build twopointers_teleport

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

package medium

// @lc code=start
// Create a map to store the last position of dupe
type CharSetWithPosition map[rune]int

func (s *CharSetWithPosition) Add(c rune, pos int) {
	(*s)[c] = pos
}

func (s *CharSetWithPosition) GetPos(c rune) int {
	if s.Contains(c) {
		return (*s)[c]
	}

	return -1
}

func (s *CharSetWithPosition) Contains(c rune) bool {
	_, exists := (*s)[c]
	return exists
}

func lengthOfLongestSubstring(s string) int {
	left := 0      // pointer refering to the start of window
	maxLength := 0 // known max length from the window

	// map to track the position of dupe
	// will be used to teleport the
	charWindow := make(CharSetWithPosition)

	for right, char := range s {
		// before we can move the right pointer
		// we need to know if the character in right pointer is dupe or not
		// this is to ensure that our right movement is valid
		// if it's already in the window
		// move the start of window to next of dupe position (dupe+1)
		// this way, we won't need to traverse the chracter one by one
		// make sure to use the largest most left position
		// to avoid situation where dupe char is positioned before left pointer
		// e.g. abbadjf -> first a is before b, and found another a after dupe b
		if charWindow.Contains(char) {
			left = max(left, charWindow.GetPos(char)+1)
		}

		// add position of current char
		// and calculate the max window size
		charWindow.Add(char, right)
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

// @lc code=end
