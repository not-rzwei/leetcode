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
	// initialize queue with root node
	queue := []*TreeNode{
		root,
	}

	for len(queue) > 0 {
		// shift the node from queue
		node := queue[0]
		queue = queue[1:len(queue)]

		if node == nil {
			continue
		}

		// swap node children
		node.Right, node.Left = node.Left, node.Right

		// add node children to the queue to traverse siblings
		queue = append(queue, node.Right)
		queue = append(queue, node.Left)
	}

	return root
}

// @lc code=end
