
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode = nil
	for curr := head; curr != nil; curr = curr.Next {
		// prepend to newHead
		newHead = &ListNode{Val: curr.Val, Next: newHead}
	}

	return newHead
}