//go:build prepostfix_inplace
// +build prepostfix_inplace

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

package medium

import "fmt"

// @lc code=start

/**
build prefix on result
result[i] = prefix[i-1] x nums[i]
nums     1 | 2 | 3
result   1 | 1 | 2

build postfix on result
result[i] = postfix[i+1] x nums[i]
nums     1 | 2 | 3
result   6 | 3 | 2

**/

func productExceptSelf(nums []int) (result []int) {
	n := len(nums)
	result = make([]int, n)

	// calculate prefix product
	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	// calculate postfix * prefix
	postfixSum := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= postfixSum
		postfixSum *= nums[i]
	}

	fmt.Println(result)
	return
}

// @lc code=end
