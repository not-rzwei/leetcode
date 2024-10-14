//go:build dfs
// +build dfs

/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 */

package easy

import (
	. "leetcode/helper"
)

// @lc code=start
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	maxDepth(root, &maxDiameter)

	return maxDiameter
}

func maxDepth(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left, maxDiameter)
	rightDepth := maxDepth(root.Right, maxDiameter)

	*maxDiameter = max(*maxDiameter, leftDepth+rightDepth)

	return 1 + max(leftDepth, rightDepth)
}

// @lc code=end
