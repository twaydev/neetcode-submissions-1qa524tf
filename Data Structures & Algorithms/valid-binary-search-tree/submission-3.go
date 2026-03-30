
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
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr := node.node

		if curr.Val >= node.max || curr.Val <= node.min {
			return false
		}

		// push to continue processing
		if curr.Right != nil {
			stack = append(stack, &StackNode{node: curr.Right, min: curr.Val, max: node.max})
		}
		if curr.Left != nil {
			stack = append(stack, &StackNode{node: curr.Left, min: node.min, max: curr.Val})
		}

	}
	return true
}