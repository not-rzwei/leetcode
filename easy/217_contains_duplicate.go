//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */
package easy

// @lc code=start
func containsDuplicate(nums []int) bool {
	visited := make(map[int]bool)
	for _, num := range nums {
		if _, isVisited := visited[num]; isVisited {
			return true
		}

		visited[num] = true
	}

	return false
}

// @lc code=end
