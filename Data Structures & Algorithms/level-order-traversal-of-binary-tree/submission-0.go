/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    
	if root == nil {
		return nil
	}

	res := [][]int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		tmpLevelVal := []int{}
		for i := len(stack); i > 0; i-- {
			// shift node from FIFO
			curr := stack[0]
			stack = stack[1:]

			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}

			tmpLevelVal = append(tmpLevelVal, curr.Val)
		}
		res = append(res, tmpLevelVal)
	}

	return res
}
