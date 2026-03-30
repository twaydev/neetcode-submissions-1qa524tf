func mergeTwoLists(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	mergedHead := &ListNode{} // must create a first node to use .Next
	curr := mergedHead
	for head1 != nil && head2 != nil {

		if head1.Val < head2.Val {
			curr.Next = head1
			head1 = head1.Next
		} else {
			curr.Next = head2
			head2 = head2.Next
		}
		curr = curr.Next
	}
	curr.Next = head1
	if head1 == nil {
		curr.Next = head2
	}

	return mergedHead.Next // remove the first node
}