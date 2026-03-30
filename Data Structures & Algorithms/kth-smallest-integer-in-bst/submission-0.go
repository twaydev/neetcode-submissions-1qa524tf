/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    
	stack := []*TreeNode{}
	count := 1
	curr := root
	for curr != nil || len(stack) > 0 {
		// push all left nodes to stack
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// pop to process
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// fmt.Printf("%v - %v\n", curr.Val, count)
		if count == k {
			return curr.Val
		}
		// push all right nodes to stack
		curr = curr.Right
		count++
	}
	return -1
}
