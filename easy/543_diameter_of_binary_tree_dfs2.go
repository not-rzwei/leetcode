//go:build dfs2
// +build dfs2

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
	if root == nil {
		return 0
	}

	return max(
		diameterOfBinaryTree(root.Left),
		diameterOfBinaryTree(root.Right),
		maxDepth(root.Left)+maxDepth(root.Right),
	)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// @lc code=end
