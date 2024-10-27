//go:build bruteforce
// +build bruteforce

package medium

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
// will timeout
func productExceptSelf(nums []int) (result []int) {
	for i := 0; i < len(nums); i++ {
		product := 1 // initialize product
		for j := 0; j < len(nums); j++ {
			if i == j { // skip if pointers are the same position
				continue
			}
			product *= nums[j]
		}
		result = append(result, product)
	}

	return
}

// @lc code=end
