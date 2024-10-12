//go:build ignore
// +build ignore

/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
 */

package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func invertTree(root *TreeNode) *TreeNode {
	// base case
	// return root if it's nil
	if root == nil {
		return root
	}

	// recursive case
	// swap children recursively
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// @lc code=end
