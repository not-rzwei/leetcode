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
type NodeWithDepth struct {
	Node  *TreeNode
	Depth int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Initialize the depth with 0
	// And starts the queue with root node
	depth := 0
	queue := []*NodeWithDepth{
		{
			Node:  root,
			Depth: 1,
		},
	}

	// Loop until queue is empty
	for len(queue) > 0 {
		// Shift the node from stack
		node := queue[0]
		queue = queue[1:len(queue)]

		// Stop checking if it's node from a leaf
		// Meaning it's nil
		if node.Node == nil {
			continue
		}

		// Compare known depth with current depth
		// Set to whichever deeper
		depth = max(depth, node.Depth)

		// Add node's children to stack regardless nil or not
		// We have nil checking above
		queue = append(queue, &NodeWithDepth{
			Node:  node.Node.Left,
			Depth: 1 + node.Depth,
		})

		queue = append(queue, &NodeWithDepth{
			Node:  node.Node.Right,
			Depth: 1 + node.Depth,
		})
	}

	return depth
}

// @lc code=end
