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
	// And starts the stack with root node
	depth := 0
	stack := []*NodeWithDepth{
		{
			Node:  root,
			Depth: 1,
		},
	}

	// Loop until stack is empty
	for len(stack) > 0 {
		// Pop the node from stack
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

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
		stack = append(stack, &NodeWithDepth{
			Node:  node.Node.Left,
			Depth: 1 + node.Depth,
		})

		stack = append(stack, &NodeWithDepth{
			Node:  node.Node.Right,
			Depth: 1 + node.Depth,
		})
	}

	return depth
}

// @lc code=end
