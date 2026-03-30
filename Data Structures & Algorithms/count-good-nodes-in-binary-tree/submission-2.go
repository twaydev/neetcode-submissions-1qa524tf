/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type QueueNode struct {
	node   *TreeNode
	maxVal int
}

func goodNodes(root *TreeNode) int {
    
	if root.Left == nil && root.Right == nil {
		return 1
	}

	goodNodes := 0
	stack := []*QueueNode{&QueueNode{node: root, maxVal: -101}}
	for len(stack) > 0 {
		// for i := len(stack); i > 0; i-- {
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if curr.node.Val >= curr.maxVal {
				goodNodes++
				curr.maxVal = curr.node.Val
			}

			if curr.node.Right != nil {
				stack = append(stack, &QueueNode{node: curr.node.Right, maxVal: curr.maxVal})
			}
			if curr.node.Left != nil {
				stack = append(stack, &QueueNode{node: curr.node.Left, maxVal: curr.maxVal})
			}
		// }
	}

	return goodNodes
}
