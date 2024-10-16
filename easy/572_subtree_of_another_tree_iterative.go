//go:build iterative
// +build iterative

/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
 */
package easy

import . "leetcode/helper"

// @lc code=start

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	stack := []*TreeNode{root} // stack to traverse preorderly
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// if node is nil, continue the loop
		// since no checking adding nil children
		if pop == nil {
			continue
		}

		// check if subRoot is the same tree to current node
		if isSubtreeIsSameTree(pop, subRoot) {
			return true
		}

		// traverse left and right
		stack = append(stack, pop.Left)
		stack = append(stack, pop.Right)
	}

	return false
}

func isSubtreeIsSameTree(p *TreeNode, q *TreeNode) bool {
	stack := [][]*TreeNode{{p, q}}

	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		np, nq := pop[0], pop[1]
		stack = stack[:len(stack)-1]

		// if both nodes are nil, continue the loop
		// since no checking adding nil children
		if np == nil && nq == nil {
			continue
		}

		// return false if value are not the same or one of them is nil
		if np == nil || nq == nil || np.Val != nq.Val {
			return false
		}

		// traverse left and right
		stack = append(stack, []*TreeNode{np.Left, nq.Left})
		stack = append(stack, []*TreeNode{np.Right, nq.Right})
	}

	return true
}

// @lc code=end
