//go:build twopointers
// +build twopointers

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

package easy

// @lc code=start
func removeDuplicates(nums []int) int {
	// slow pointer to track the unique element
	// fast pointer to check the unique element
	// fast pointer will lopop through the numbers
	slow, fast := 0, 0

	for fast < len(nums) {
		// if slow and fast value are different (unique, not a dupe)
		// move slow pointer and replace the value with fast value
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}

		// move forward the fast pointer
		fast++
	}

	// return the slow (that tracks the unique element)
	// +1 because slow starts from 0
	return slow + 1
}

// @lc code=end
