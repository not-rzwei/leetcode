//go:build prepostfix
// +build prepostfix

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

package medium

// @lc code=start

/**
index    0   1   2   3
nums     1 | 2 | 3 |=> 0-index
prefix   1 | 1 | 2 | 6 => 1-index
postfix  6 | 6 | 3 | 3 => 1-index

result[i] = prefix[i] + postfix[i+1] => 0-index
result     | 6 | 3 | 2
**/

func productExceptSelf(nums []int) (result []int) {
	n := len(nums)

	// create prefix and initialize the first index to 1
	prefix := make([]int, n+1)
	prefix[0] = 1
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] * nums[i]
	}

	// create postfix and initialize the first index to 1
	postfix := make([]int, n+1)
	postfix[n] = 1
	for i := n - 1; i >= 0; i-- {
		postfix[i] = postfix[i+1] * nums[i]
	}

	// iterate tru nums
	// perform multiplication on prefix[i] and postfix[i+1]
	for i := 0; i < n; i++ {
		result = append(result, prefix[i]*postfix[i+1])
	}

	return
}

// @lc code=end
