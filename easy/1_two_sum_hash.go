//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package easy

// @lc code=start

func twoSum(nums []int, target int) []int {
	visited := make(map[int]int, len(nums))

	for index, num := range nums {
		diff := target - num
		if vIndex, isVisited := visited[diff]; isVisited {
			return []int{vIndex, index}
		}

		visited[num] = index
	}

	return nil
}

// @lc code=end
