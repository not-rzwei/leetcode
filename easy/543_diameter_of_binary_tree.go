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
	// track the maximum diameter
	_, maxDiameter := maxDepth(root)

	return maxDiameter
}

// return depth and diameter
func maxDepth(root *TreeNode) (int, int) {
	// base case
	// 0 depth for nil node
	if root == nil {
		return 0, 0
	}

	// recursive case
	// post order traversal
	// calculating the depth of left nodes first
	// then right nodes
	leftDepth, leftDiameter := maxDepth(root.Left)
	rightDepth, rightDiameter := maxDepth(root.Right)

	// process the root node (parent node)
	// calculate the max diameter
	// piggyback max depth calculation
	maxDiameter := max(leftDiameter, rightDiameter, leftDepth+rightDepth)

	// return the depth of node
	return 1 + max(leftDepth, rightDepth), maxDiameter
}

// @lc code=end
