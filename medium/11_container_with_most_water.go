/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

package medium

// @lc code=start
func maxArea(height []int) int {
	// initialize left and right pointer
	// left will the the start of container
	// right will be the end of container
	// set right right most to find the largest area first
	// then shrink down the length (left or right pointer)
	left, right := 0, len(height)-1
	maxArea := 0

	// loop until no more length to shrink
	for left < right {
		// set the min height of pointers
		// we do this to avoid overflowing water
		// if we chose the highest height
		// water will overflow
		minHeight := min(height[left], height[right])
		curArea := minHeight * (right - left)
		maxArea = max(maxArea, curArea)

		// once calculate the max area
		// shrink the left or right pointer
		// shrink the lower height
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// @lc code=end
