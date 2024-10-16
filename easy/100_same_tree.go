//go:build recursive
// +build recursive

/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 */

package easy

import . "leetcode/helper"

// @lc code=start
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// true if both node is nil
	// meaning traversal can reach leaf node
	if p == nil && q == nil {
		return true
	}

	// false if one of them is nil
	// meaning one of them has different value
	if p == nil || q == nil {
		return false
	}

	// false if one of them has different value
	if p.Val != q.Val {
		return false
	}

	// traverse preorderly
	isSameLeftNode := isSameTree(p.Left, q.Left)
	isSameRightNode := isSameTree(p.Right, q.Right)

	return isSameLeftNode && isSameRightNode
}

// @lc code=end
