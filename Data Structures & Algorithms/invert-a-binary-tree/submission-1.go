/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    
	if root == nil {
		return root
	}

	// stack-based itegration: lấy giá trị trong stack ra để xử lý
	// declare stack with root at the first node
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]  // lấy giá trị cuối slice
		stack = stack[:len(stack)-1] // bỏ giá trị "node" ra khỏi slice

		// invert tree
		dummyNode := node.Left
		node.Left = node.Right
		node.Right = dummyNode

		// đẩy giá trị left right của node hiện tại vào để xử lý tiếp
		// pre-order nên sẽ đẩy right và trước để left được xử lý trước
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return root
}
