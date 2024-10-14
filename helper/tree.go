package helper

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildBTree(values []interface{}, index int) *TreeNode {
	if index >= len(values) || values[index] == nil {
		return nil
	}

	node := &TreeNode{
		Val: values[index].(int),
	}
	node.Left = BuildBTree(values, 2*index+1)  // Left child index
	node.Right = BuildBTree(values, 2*index+2) // Right child index

	return node
}
