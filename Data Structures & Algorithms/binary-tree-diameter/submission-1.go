/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type NodeState struct {
	node  *TreeNode
	state int // 0: PRE 1:POST
}
func diameterOfBinaryTree(root *TreeNode) int {
    if root.Left == nil && root.Right == nil {
		return 0
	}

	// declare maxDiameter
	maxDiameter := 0
	// declare stack as FILO
	stack := []*NodeState{{node: root, state: 0}}
	// declare nodeHeights
	heights := make(map[*TreeNode]int)

	// bottom-up pattern (post-order DFS)
	// we will calculate height of leaves then store to heights map
	// the parents will use values from heights map to calculate height itself
	for len(stack) > 0 {
		// get tail as current node
		curr := stack[len(stack)-1]

		if curr.state == 0 {
			curr.state = 1
			if curr.node.Right != nil {
				stack = append(stack, &NodeState{node: curr.node.Right, state: 0})
			}
			if curr.node.Left != nil {
				stack = append(stack, &NodeState{node: curr.node.Left, state: 0})
			}
		} else {
			// POP node to process
			stack = stack[:len(stack)-1]

			// these lines are only run with parent.
			leftHeight := heights[curr.node.Left]
			rightHeight := heights[curr.node.Right]
			// store height of parent node to hashmap
			heights[curr.node] = 1 + max(leftHeight, rightHeight)

			// update maxDiameter
			maxDiameter = max(maxDiameter, leftHeight+rightHeight)
		}
	}

	return maxDiameter
}
