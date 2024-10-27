/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */

package medium

// @lc code=start
func twoSum(numbers []int, target int) []int {
	// start from start and end position
	left, right := 0, len(numbers)-1

	// loop until left meet right
	for left < right {
		sum := numbers[left] + numbers[right]

		// if sum is match with target
		// return that position
		if sum == target {
			return []int{left + 1, right + 1}
		}

		// if bigger, move the right pointer backward
		if sum > target {
			right--
		} else { // otherwise move left pointer forward
			left++
		}
	}

	return []int{}
}

// @lc code=end
