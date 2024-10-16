//go:build recursive
// +build recursive

/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
 */
package easy

import . "leetcode/helper"

// @lc code=start

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	// check sameTree for every node of root
	if sameTree(root, subRoot) {
		return true
	}

	// traverse left and right
	// return true if found same tree in either left or right nodes
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func sameTree(p *TreeNode, q *TreeNode) bool {
	// if both nodes are nil, return true
	// if one of them reach nil return false
	if p == nil || q == nil {
		return p == q
	}

	// if value are not the same, return false
	if p.Val != q.Val {
		return false
	}

	// traverse left and right
	// only true if both left and right are the same
	return sameTree(p.Left, q.Left) && sameTree(p.Right, q.Right)
}

// @lc code=end
