//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */
package easy

// @lc code=start
func search(nums []int, target int) int {
	// pointer to define the search boundary
	// starting from left ending with right
	left, right := 0, len(nums)-1

	// loop until boundary become 1
	// left pointer and right pointer meet each other
	for left <= right {
		// get the middle of boundary
		// return the position if target is in the middle
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		// if middle number is bigger than target
		// reduce the boundary
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

// @lc code=end
