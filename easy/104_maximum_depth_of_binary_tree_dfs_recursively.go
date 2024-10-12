//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */

package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
// Create custom node to track its depth
func maxDepth(root *TreeNode) int {
	// base case
	// return 0 when it hits the leaf node
	if root == nil {
		return 0
	}

	// add 1 to the max depth of children recursively
	// 1 because the node itself counts as a level
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// @lc code=end
