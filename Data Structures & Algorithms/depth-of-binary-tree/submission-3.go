/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	// bfs - breadth first search
	stack := []*TreeNode{root}
	lvl := 0
	for len(stack) > 0 {
		lvl++
		for i := len(stack); i > 0; i-- {
			// pop node
			node := stack[0]
			stack = stack[1:]

			// insert node
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	return lvl
}
