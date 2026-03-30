
type StackNode struct {
	node *TreeNode
	min  int
	max  int
}

func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	stack := []*StackNode{&StackNode{node: root, min: -1001, max: 1001}}
	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]

		if curr.node.Val >= curr.max || curr.node.Val <= curr.min {
			return false
		}

		// push to continue processing
		if curr.node.Right != nil {
			stack = append(stack, &StackNode{node: curr.node.Right, min: curr.node.Val, max: curr.max})
		}
		if curr.node.Left != nil {
			stack = append(stack, &StackNode{node: curr.node.Left, min: curr.min, max: curr.node.Val})
		}

	}
	return true
}