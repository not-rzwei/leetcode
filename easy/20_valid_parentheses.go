//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

package easy

// @lc code=start
func isValid(s string) bool {
	// stack to track the opening parenthesis
	stack := []rune{}

	// map opening to closing
	// this is to match opening in stack
	pairs := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	// loop through the string
	for _, c := range s {
		// add opening to stack
		if _, isOpening := pairs[c]; isOpening {
			stack = append(stack, c)
			continue
		}

		// return false if stack is empty
		// meaning no openings but found a closing
		if len(stack) == 0 {
			return false
		}

		// pop the opening from stack
		opening := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// return false if current closing doesn't match with opening pair
		if pairs[opening] != c {
			return false
		}
	}

	return len(stack) == 0
}

// @lc code=end
