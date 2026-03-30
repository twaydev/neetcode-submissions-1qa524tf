/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		lvl := []int{}
		for i := len(stack); i > 0; i-- {

			curr := stack[0]
			stack = stack[1:]

			lvl = append(lvl, curr.Val)

			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
		}
		// pop to get right node
		res = append(res, lvl[len(lvl)-1])
	}
	return res
}
