//go:build iterative
// +build iterative

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

	stack := []*TreeNode{root}          // stack to do post order traversal
	depth := make(map[*TreeNode]int, 0) // track the depth of each node
	diameter := 0                       // max known diameter

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
		depth[node] = 1 + max(depth[node.Left], depth[node.Right])   // depth = 1 + max(depth(L), (R))
		diameter = max(diameter, depth[node.Left]+depth[node.Right]) // diameter = depth(L) + depth(R)
		stack = stack[:len(stack)-1]                                 // remove from the stack
	}

	return diameter
}

// @lc code=end
