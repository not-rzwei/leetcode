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
	// initialize stack with root node
	stack := []*TreeNode{
		root,
	}

	// loop until stack is empty
	for len(stack) > 0 {
		// pop the node
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// skip if node is empty
		if node == nil {
			continue
		}

		// swap node children
		node.Right, node.Left = node.Left, node.Right

		// add node children to the stack to traverse the node children
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}

	return root
}

// @lc code=end
