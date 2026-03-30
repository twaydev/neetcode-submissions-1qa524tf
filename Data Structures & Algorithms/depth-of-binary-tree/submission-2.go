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


	stack := make([][]interface{}, 0)
	stack = append(stack, []interface{}{root, 1})
	maxDepth := 0
	for len(stack) > 0 {

		// pop giá trị trong stack làm current Node và current Depth
		node := stack[len(stack)-1][0].(*TreeNode)
		depth := stack[len(stack)-1][1].(int)
		stack = stack[:len(stack)-1]

		// đẩy giá trị left right của node hiện tại vào để xử lý tiếp
		// pre-order nên sẽ đẩy right và trước để left được xử lý trước
		if node != nil {
			stack = append(stack, []interface{}{node.Right, depth + 1})
			stack = append(stack, []interface{}{node.Left, depth + 1})
			maxDepth = max(maxDepth, depth)
		}

	}

	return maxDepth
}
