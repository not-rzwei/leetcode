//go:build iterative
// +build iterative

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */
package easy

import . "leetcode/helper"

// @lc code=start
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := []*TreeNode{root}          // stack to do post order traversal
	depth := make(map[*TreeNode]int, 0) // track the depth of each node
	balanced := true                    // balanced status

	// traverse post orderly
	for len(stack) > 0 {
		node := stack[len(stack)-1]

		// go left until hit leaf node
		if node.Left != nil && depth[node.Left] == 0 {
			stack = append(stack, node.Left)
			continue
		}

		// go right until hit leaf node
		if node.Right != nil && depth[node.Right] == 0 {
			stack = append(stack, node.Right)
			continue
		}

		// process the node
		leftDepth := depth[node.Left]
		rightDepth := depth[node.Right]
		depth[node] = 1 + max(leftDepth, rightDepth)

		// diff left depth and right depth
		// is balanced if diff <= 1
		depthDiff := max(leftDepth, rightDepth) - min(leftDepth, rightDepth)
		balanced = balanced && depthDiff <= 1
		stack = stack[:len(stack)-1]
	}

	return balanced
}

// @lc code=end
