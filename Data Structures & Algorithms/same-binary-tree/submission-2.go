/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    
	if p == nil && q == nil {
		return true
	}
	// if p == nil || q == nil {
	// 	return false
	// }
	stack := [][]*TreeNode{}
	stack = append(stack, []*TreeNode{p, q})
	for len(stack) > 0 {
		// pop node to process
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Printf("%v - %v\n", curr[0], curr[1])

		if curr[0] == nil && curr[1] == nil {
			continue
		}
		if curr[0] == nil || curr[1] == nil {
			return false
		}
		if curr[0].Val != curr[1].Val {
			return false
		}
		// if curr[0].Right != nil || curr[1].Right != nil {
		// 	return false
		// }
		// if curr[0].Left != nil || curr[1].Left != nil {
		// 	return false
		// }

		// push to process
		// if curr[0].Right != nil && curr[1].Right != nil {
		stack = append(stack, []*TreeNode{curr[0].Right, curr[1].Right})
		// }

		// if curr[0].Left != nil && curr[1].Left != nil {
		stack = append(stack, []*TreeNode{curr[0].Left, curr[1].Left})
		// }

	}

	return true
}
