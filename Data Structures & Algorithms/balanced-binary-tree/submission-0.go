
type NodeState struct {
	node  *TreeNode
	state int // 0:PRE 1:POST
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isBalanced := true
	stack := []*NodeState{{node: root, state: 0}}
	nodeHeights := make(map[*TreeNode]int)
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		if curr.state == 0 {
			// push to stack
			curr.state = 1
			if curr.node.Right != nil {
				stack = append(stack, &NodeState{node: curr.node.Right, state: 0})
			}
			if curr.node.Left != nil {
				stack = append(stack, &NodeState{node: curr.node.Left, state: 0})
			}
		} else {
			// pop to process
			stack = stack[:len(stack)-1]
			leftH := nodeHeights[curr.node.Left]
			rightH := nodeHeights[curr.node.Right]
			if leftH-rightH > 1 || rightH-leftH > 1 {
				isBalanced = false
			}
			nodeHeights[curr.node] = max(leftH, rightH) + 1
		}
	}

	return isBalanced
}
