//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

package easy

// @lc code=start
// func to check whether the byte is within alphanumeric bound
func isAlphanumeric(b byte) bool {
	return (b >= 'A' && b <= 'Z') ||
		(b >= 'a' && b <= 'z') ||
		(b >= '0' && b <= '9')
}

func toLower(b byte) byte {
	// only do it for uppercase characrter
	if b >= 'A' && b <= 'Z' {
		// a - A return ASCII offset
		// lowercase has higher codepoint
		return b + ('a' - 'A')
	}

	return b
}

func isPalindrome(s string) bool {
	// pointers to compare both end
	left, right := 0, len(s)-1

	// loop until both pointer adjacent to each other
	for left < right {
		// move forward left pointer until found alphanumeric character
		// and only move if it has right
		for !isAlphanumeric(s[left]) && left < right {
			left++
		}

		// move backward right pointer until found alphanumeric character
		for !isAlphanumeric(s[right]) && right > left {
			right--
		}

		// compare both pointer, return false if not the same
		// comparison is case insensitive
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		// move both pointers
		// if not, it'll loop endlessy
		left++
		right--
	}

	return true
}

// @lc code=end
