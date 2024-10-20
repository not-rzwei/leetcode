//go:build bruteforce
// +build bruteforce

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package easy

// @lc code=start

func twoSum(nums []int, target int) []int {
	for i := 0; i <= len(nums)-1; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			diff := nums[i] + nums[j]
			if diff == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// @lc code=end
