//go:build recursive
// +build recursive

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */
package easy

import . "leetcode/helper"

// @lc code=start
func isBalanced(root *TreeNode) bool {
	_, balanced := isDepthBalanced(root)

	return balanced
}

func isDepthBalanced(root *TreeNode) (int, bool) {
	// base case
	// nil depth is 0
	if root == nil {
		return 0, true
	}

	// postorder traversal
	// get the left depth first
	// then the right depth
	leftDepth, _ := isDepthBalanced(root.Left)
	rightDepth, _ := isDepthBalanced(root.Right)

	// process the node
	// compare the diff, not balanced if more than 1
	depthDiff := max(leftDepth, rightDepth) - min(leftDepth, rightDepth)
	bothBalanced := depthDiff <= 1

	// return depth and balanced status
	return 1 + max(leftDepth, rightDepth), bothBalanced
}

// @lc code=end
