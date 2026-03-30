/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
	// reverse the origin list
	var reverseNode *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = reverseNode
		reverseNode = head

		// fmt.Printf("reverse: %v\n", head.Val)
		head = tmp
	}

	var result *ListNode
	head = reverseNode
	count := 1
	for head != nil {
		// fmt.Printf("count: %v, head: %v\n", count, head.Val)
		if count != n {
			result = &ListNode{Val: head.Val, Next: result}
			// fmt.Printf("add: %v\n", head.Val)
		}
		count++
		head = head.Next
	}

	return result
}
