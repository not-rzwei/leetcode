//go:build iterative
// +build iterative

/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 */

package easy

import . "leetcode/helper"

// @lc code=start
func isSameTree(p *TreeNode, q *TreeNode) bool {
	stack := [][]*TreeNode{{p, q}} // stack to traverse, assign both nodes

	// traverse preorderly
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		np, nq := pop[0], pop[1]

		// if both reach leaf, continue to other nodes
		if np == nil && nq == nil {
			continue
		}

		// if one of them reach leaf first, abort
		if np == nil || nq == nil {
			return false
		}

		// abort if they have different value
		if np.Val != nq.Val {
			return false
		}

		// traverse left and right
		stack = append(stack, []*TreeNode{np.Left, nq.Left})
		stack = append(stack, []*TreeNode{np.Right, nq.Right})
	}

	// return true because both trees have been traversed
	return true
}

// @lc code=end
