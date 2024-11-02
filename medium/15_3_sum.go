/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

package medium

import "slices"

// @lc code=start
func threeSum(nums []int) (result [][]int) {
	// sort the nums to make it monotonic
	// so w can use binary search
	slices.Sort(nums)

	n := len(nums) // shorthand

	// loop through the nums
	// start with i pointer for first item in triplet
	for i := 0; i < n; i++ {
		// break if first item in triplet is higher than 0
		// meaning, second and third item can't sum to 0
		if nums[i] > 0 {
			break
		}

		// skip if current element is same with previous element
		// since we can't have duplicate triplets
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// binary search
		// left for second item in tripet
		// right for third item in triplet
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			// if items in triplet sum to 0
			// append the triplet to result
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// remember we're finding all possible triplets
				// move left to perform another binary search
				// and move pointer to the right until no more dupe number in L
				// e.g. 1 1 2  1 1 2
				//      L -->      L
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if sum > 0 { // if sum is higher, reduce the right side
				right--
			} else { // otherwise reduce the left side
				left++
			}
		}
	}

	return
}

// @lc code=end
